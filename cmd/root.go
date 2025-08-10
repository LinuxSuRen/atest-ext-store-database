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
package cmd

import (
	"fmt"
	"net"

	"github.com/linuxsuren/api-testing/pkg/testing/remote"
	"github.com/linuxsuren/atest-ext-store-database/pkg"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func NewRootCmd() (cmd *cobra.Command) {
	opt := &option{}
	cmd = &cobra.Command{
		Use:  "atest-store-database",
		RunE: opt.runE,
	}
	flags := cmd.Flags()
	flags.IntVarP(&opt.port, "port", "p", 7073, "The port of gRPC server")
	return
}

func (o *option) runE(cmd *cobra.Command, args []string) (err error) {
	var removeServer pkg.RemoteServer
	if removeServer, err = pkg.NewRemoteServer(); err != nil {
		return
	}

	var lis net.Listener
	lis, err = net.Listen("tcp", fmt.Sprintf(":%d", o.port))
	if err != nil {
		return
	}

	gRPCServer := grpc.NewServer()
	remote.RegisterLoaderServer(gRPCServer, removeServer)
	cmd.Println("Data extension is running at port", o.port)

	go func() {
		<-cmd.Context().Done()
		gRPCServer.Stop()
	}()

	err = gRPCServer.Serve(lis)
	return
}

type option struct {
	port int
}
