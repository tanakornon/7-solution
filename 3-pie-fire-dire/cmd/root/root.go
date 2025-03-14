package root

import (
	"log"
	"os"
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
			if err := protocol.ServeGRPC(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		},
	}

	serveRESTCmd := &cobra.Command{
		Use:   "rest",
		Short: "Start REST server",
		Run: func(cmd *cobra.Command, args []string) {
			if err := protocol.ServeREST(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		},
	}

	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start both gRPC and REST servers concurrently",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				if err := protocol.ServeGRPC(); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}()

			go func() {
				if err := protocol.ServeREST(); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}()

			select {}
		},
	}

	rootCmd.AddCommand(
		serveGRPCCmd,
		serveRESTCmd,
		serveCmd,
	)

	cobra.CheckErr(rootCmd.Execute())
}
