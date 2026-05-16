package cmd

import (
	"github.com/ramadhanalfarisi/go-codebase/app/grpc"
	"github.com/spf13/cobra"
)

var GrpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start the GRPC server",
	Long:  `Start the GRPC server with all the necessary configurations and dependencies`,
	Run: func(cmd *cobra.Command, args []string) {
		g := grpc.NewGrpc()
		g.Run()
	}}
