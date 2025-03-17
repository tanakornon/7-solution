package root

import (
	"pie-fire-dire/cmd/protocol"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use: "pie-fire-dire",
	}

	serveGRPCCmd := &cobra.Command{
		Use:   "grpc",
		Short: "Start gRPC server",
		Run: func(cmd *cobra.Command, args []string) {
			protocol.NewServer().WithREST().Start()
		},
	}

	serveRESTCmd := &cobra.Command{
		Use:   "rest",
		Short: "Start REST server",
		Run: func(cmd *cobra.Command, args []string) {
			protocol.NewServer().WithGRPC().Start()
		},
	}

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start both gRPC and REST servers concurrently",
		Run: func(cmd *cobra.Command, args []string) {
			protocol.NewServer().WithGRPC().WithREST().Start()
		},
	}

	rootCmd.AddCommand(
		serveGRPCCmd,
		serveRESTCmd,
		serveCmd,
	)

	cobra.CheckErr(rootCmd.Execute())
}
