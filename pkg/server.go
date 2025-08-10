/*
Copyright 2025 API Testing Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package pkg

import (
	"context"

	"github.com/linuxsuren/atest-ext-store-database/ui"

	"github.com/linuxsuren/api-testing/pkg/server"
	"github.com/linuxsuren/api-testing/pkg/testing/remote"
)

type databaseExtension struct {
	remote.UnimplementedLoaderServer
}

type RemoteServer interface {
	remote.LoaderServer
	UIExtension
}

type UIExtension interface {
	GetPageOfJS(ctx context.Context, in *server.SimpleName) (*server.CommonResult, error)
	GetPageOfCSS(ctx context.Context, in *server.SimpleName) (*server.CommonResult, error)
	GetMenus(ctx context.Context, empty *server.Empty) (*server.MenuList, error)
}

func NewRemoteServer() (server RemoteServer, err error) {
	server = &databaseExtension{}
	return
}

func (s *databaseExtension) GetMenus(ctx context.Context, empty *server.Empty) (reply *server.MenuList, err error) {
	reply = &server.MenuList{
		Data: []*server.Menu{
			{
				Name:  "database",
				Index: "database",
				Icon:  "DataAnalysis",
			},
		},
	}
	return
}

func (s *databaseExtension) GetPageOfJS(ctx context.Context, in *server.SimpleName) (reply *server.CommonResult, err error) {
	reply = &server.CommonResult{
		Success: true,
		Message: ui.GetJS(),
	}
	return
}

func (s *databaseExtension) GetPageOfCSS(ctx context.Context, in *server.SimpleName) (reply *server.CommonResult, err error) {
	reply = &server.CommonResult{
		Success: true,
		Message: ui.GetCSS(),
	}
	return
}
