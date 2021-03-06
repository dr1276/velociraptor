package reporting

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/alexmullins/zip"
	"github.com/pkg/errors"
	actions_proto "www.velocidex.com/golang/velociraptor/actions/proto"
	config_proto "www.velocidex.com/golang/velociraptor/config/proto"
	"www.velocidex.com/golang/velociraptor/file_store/csv"
	"www.velocidex.com/golang/velociraptor/uploads"
	"www.velocidex.com/golang/velociraptor/utils"
	"www.velocidex.com/golang/vfilter"
)

type Container struct {
	sync.Mutex // Serialize access to the zip file.
	fd         io.WriteCloser
	zip        *zip.Writer

	Password     string
	delegate_zip *zip.Writer
}

func (self *Container) StoreArtifact(
	config_obj *config_proto.Config,
	ctx context.Context,
	scope *vfilter.Scope,
	vql *vfilter.VQL,
	query *actions_proto.VQLRequest,
	format string) error {

	// Dont store un-named queries but run them anyway.
	if query.Name == "" {
		for _ = range vql.Eval(ctx, scope) {
		}
		return nil
	}

	// Queue the query into a temp file in order to allow
	// interleaved uploaded files to be written to the zip
	// file. Zip files only support a single writer at a time.
	tmpfile, err := ioutil.TempFile("", "vel")
	if err != nil {
		return errors.WithStack(err)
	}

	defer os.Remove(tmpfile.Name()) // clean up

	var sanitized_name string

	switch format {
	case "csv", "":
		// In this instance we want to make / unescaped.
		sanitized_name = query.Name + ".csv"

		csv_writer := csv.GetCSVAppender(
			scope, tmpfile, true /* write_headers */)
		defer csv_writer.Close()

		for row := range vql.Eval(ctx, scope) {
			csv_writer.Write(row)
		}

	case "jsonl", "json":
		// In this instance we want to make / unescaped.
		sanitized_name = query.Name + ".json"

		for row := range vql.Eval(ctx, scope) {
			// Re-serialize it as compact json.
			serialized, err := json.Marshal(row)
			if err != nil {
				continue
			}

			_, err = tmpfile.Write(serialized)
			if err != nil {
				return errors.WithStack(err)
			}

			// Separate lines with \n
			tmpfile.Write([]byte("\n"))
		}

	default:
		return errors.New("Format not supported")
	}

	writer, err := self.getZipFileWriter(string(sanitized_name))
	if err != nil {
		return err
	}

	tmpfile.Seek(0, 0)

	_, err = utils.Copy(ctx, writer, tmpfile)
	return err
}

func (self *Container) getZipFileWriter(name string) (io.Writer, error) {
	if self.Password == "" {
		return self.zip.Create(string(name))
	}

	// Zip file encryption is not great because it only encrypts
	// the content of the file, and not its directory. We want to
	// do better than that - so we create another zip file inside
	// the original zip file and encrypt that.
	if self.delegate_zip == nil {
		fd, err := self.zip.Encrypt("data.zip", self.Password)
		if err != nil {
			return nil, err
		}

		self.delegate_zip = zip.NewWriter(fd)
	}

	w, err := self.delegate_zip.Create(string(name))
	return w, err
}

func (self *Container) DumpRowsIntoContainer(
	config_obj *config_proto.Config,
	output_rows []vfilter.Row,
	scope *vfilter.Scope,
	query *actions_proto.VQLRequest) error {

	if len(output_rows) == 0 {
		return nil
	}

	self.Lock()
	defer self.Unlock()

	// In this instance we want to make / unescaped.
	sanitized_name := query.Name + ".csv"
	writer, err := self.getZipFileWriter(string(sanitized_name))
	if err != nil {
		return err
	}

	csv_writer := csv.GetCSVAppender(
		scope, writer, true /* write_headers */)
	for _, row := range output_rows {
		csv_writer.Write(row)
	}

	csv_writer.Close()

	sanitized_name = query.Name + ".json"
	writer, err = self.getZipFileWriter(string(sanitized_name))
	if err != nil {
		return err
	}

	err = json.NewEncoder(writer).Encode(output_rows)
	if err != nil {
		return err
	}

	// Format the description.
	sanitized_name = query.Name + ".txt"
	writer, err = self.getZipFileWriter(string(sanitized_name))
	if err != nil {
		return err
	}

	fmt.Fprintf(writer, "# %s\n\n%s", query.Name,
		FormatDescription(config_obj, query.Description, output_rows))

	return nil
}

func sanitize(component string) string {
	component = strings.Replace(component, ":", "", -1)
	component = strings.Replace(component, "?", "", -1)
	return component
}

func (self *Container) Upload(
	ctx context.Context,
	scope *vfilter.Scope,
	filename string,
	accessor string,
	store_as_name string,
	expected_size int64,
	reader io.Reader) (*uploads.UploadResponse, error) {
	self.Lock()
	defer self.Unlock()

	var components []string
	if store_as_name == "" {
		store_as_name = filename
		components = []string{accessor}
	}

	// Normalize and clean up the path so the zip file is more
	// usable by fragile zip programs like Windows explorer.
	for _, component := range utils.SplitComponents(store_as_name) {
		if component == "." || component == ".." {
			continue
		}
		components = append(components, sanitize(component))
	}

	// Zip members must not have absolute paths.
	sanitized_name := path.Join(components...)
	writer, err := self.getZipFileWriter(sanitized_name)
	if err != nil {
		return nil, err
	}

	scope.Log("Collecting file %s (%v bytes)",
		store_as_name, expected_size)

	sha_sum := sha256.New()
	md5_sum := md5.New()

	n, err := utils.Copy(ctx, utils.NewTee(writer, sha_sum, md5_sum), reader)
	if err != nil {
		return &uploads.UploadResponse{
			Error: err.Error(),
		}, err
	}

	return &uploads.UploadResponse{
		Path:   sanitized_name,
		Size:   uint64(n),
		Sha256: hex.EncodeToString(sha_sum.Sum(nil)),
		Md5:    hex.EncodeToString(md5_sum.Sum(nil)),
	}, nil
}

func (self *Container) Close() error {
	if self.delegate_zip != nil {
		self.delegate_zip.Close()
	}

	self.zip.Close()
	return self.fd.Close()
}

func NewContainer(path string) (*Container, error) {
	fd, err := os.OpenFile(
		path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0700)
	if err != nil {
		return nil, err
	}

	zip_writer := zip.NewWriter(fd)
	return &Container{
		fd:  fd,
		zip: zip_writer,
	}, nil
}

// Turns os.Stdout into into file_store.WriteSeekCloser
type StdoutWrapper struct {
	io.Writer
}

func (self *StdoutWrapper) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (self *StdoutWrapper) Close() error {
	return nil
}

func (self *StdoutWrapper) Truncate(offset int64) error {
	return nil
}
