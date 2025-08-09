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
}

func NewRemoteServer() (server RemoteServer, err error) {
	server = &databaseExtension{}
	return
}

func (s *databaseExtension) GetMenus(ctx context.Context, empty *server.Empty) (reply *server.MenuList, err error) {
	reply = &server.MenuList{
		Data: []*server.Menu{
			{
				Name:  "data",
				Index: "data",
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
