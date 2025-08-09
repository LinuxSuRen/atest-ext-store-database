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
