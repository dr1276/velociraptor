/*
   Velociraptor - Hunting Evil
   Copyright (C) 2019 Velocidex Innovations.

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package networking

import (
	"context"

	"github.com/Velocidex/ordereddict"
	"www.velocidex.com/golang/velociraptor/glob"
	"www.velocidex.com/golang/velociraptor/uploads"
	vql_subsystem "www.velocidex.com/golang/velociraptor/vql"
	"www.velocidex.com/golang/vfilter"
)

// We also offer a VQL function to manage the upload.
// Example: select upload(file=FullPath) from glob(globs="/bin/*")

type UploadFunctionArgs struct {
	File     string `vfilter:"required,field=file,doc=The file to upload"`
	Name     string `vfilter:"optional,field=name,doc=The name of the file that should be stored on the server"`
	Accessor string `vfilter:"optional,field=accessor,doc=The accessor to use"`
}
type UploadFunction struct{}

func (self *UploadFunction) Call(ctx context.Context,
	scope *vfilter.Scope,
	args *ordereddict.Dict) vfilter.Any {

	uploader_obj, ok := scope.Resolve("$uploader")
	if !ok {
		scope.Log("upload: Uploader not configured.")
		return vfilter.Null{}
	}
	uploader, ok := uploader_obj.(uploads.Uploader)
	if ok {
		arg := &UploadFunctionArgs{}
		err := vfilter.ExtractArgs(scope, args, arg)
		if err != nil {
			scope.Log("upload: %s", err.Error())
			return vfilter.Null{}
		}

		if arg.File == "" {
			return vfilter.Null{}
		}

		err = vql_subsystem.CheckFilesystemAccess(scope, arg.Accessor)
		if err != nil {
			scope.Log("upload: %s", err)
			return vfilter.Null{}
		}

		accessor, err := glob.GetAccessor(arg.Accessor, ctx)
		if err != nil {
			scope.Log("upload: %v", err)
			return &uploads.UploadResponse{
				Error: err.Error(),
			}
		}

		file, err := accessor.Open(arg.File)
		if err != nil {
			scope.Log("upload: Unable to open %s: %s",
				arg.File, err.Error())
			return &uploads.UploadResponse{
				Error: err.Error(),
			}
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			scope.Log("upload: Unable to stat %s: %v",
				arg.File, err)
		} else if !stat.IsDir() {
			upload_response, err := uploader.Upload(
				ctx, scope, arg.File,
				arg.Accessor,
				arg.Name,
				stat.Size(), // Expected size.
				file)
			if err != nil {
				return &uploads.UploadResponse{
					Error: err.Error(),
				}
			}
			return upload_response
		}
	}
	return vfilter.Null{}
}

func (self UploadFunction) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.FunctionInfo {
	return &vfilter.FunctionInfo{
		Name: "upload",
		Doc: "Upload a file to the upload service. For a Velociraptor " +
			"client this will upload the file into the flow and store " +
			"it in the server's file store.",
		ArgType: type_map.AddType(scope, &UploadFunctionArgs{}),
	}
}

type UploadPluginArgs struct {
	Files    []string `vfilter:"required,field=files,doc=A list of files to upload"`
	Accessor string   `vfilter:"optional,field=accessor,doc=The accessor to use"`
}

type UploadPlugin struct{}

func (self *UploadPlugin) Call(
	ctx context.Context,
	scope *vfilter.Scope,
	args *ordereddict.Dict) <-chan vfilter.Row {
	output_chan := make(chan vfilter.Row)

	go func() {
		defer close(output_chan)

		arg := &UploadPluginArgs{}
		err := vfilter.ExtractArgs(scope, args, arg)
		if err != nil {
			scope.Log("upload: %s", err.Error())
			return
		}

		uploader_obj, _ := scope.Resolve("$uploader")
		uploader, ok := uploader_obj.(uploads.Uploader)
		if !ok {
			scope.Log("upload: Uploader not configured.")

			// If the uploader is not configured, we need to do
			// nothing.
			return
		}

		err = vql_subsystem.CheckFilesystemAccess(scope, arg.Accessor)
		if err != nil {
			scope.Log("upload: %s", err)
			return
		}

		accessor, err := glob.GetAccessor(arg.Accessor, ctx)
		if err != nil {
			scope.Log("upload: %v", err)
			return
		}

		for _, filename := range arg.Files {
			file, err := accessor.Open(filename)
			if err != nil {
				scope.Log("upload: Unable to open %s: %s",
					filename, err.Error())
				continue
			}

			stat, err := file.Stat()
			if err != nil {
				scope.Log("upload: Unable to stat %s: %v",
					filename, err)
			} else if !stat.IsDir() {
				upload_response, err := uploader.Upload(
					ctx, scope, filename,
					arg.Accessor,
					filename,
					stat.Size(), // Expected size.
					file)
				if err != nil {
					scope.Log("upload: Failed to upload %s: %s",
						filename, err.Error())
					continue
				}
				output_chan <- upload_response
			}
		}
	}()
	return output_chan
}

func (self UploadPlugin) Info(scope *vfilter.Scope, type_map *vfilter.TypeMap) *vfilter.PluginInfo {
	return &vfilter.PluginInfo{
		Name:    "upload",
		Doc:     "Upload files to the server.",
		RowType: type_map.AddType(scope, &uploads.UploadResponse{}),
		ArgType: type_map.AddType(scope, &UploadPluginArgs{}),
	}
}

func init() {
	vql_subsystem.RegisterFunction(&UploadFunction{})
	vql_subsystem.RegisterPlugin(&UploadPlugin{})
}
