package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"template/config"
	"template/pkg/http"
)

var serveCmd = &cobra.Command{
	Use: `serve`,
	Short: ``,
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		s := http.NewEchoServer()
		addr, err := config.C.GetString("server.http.address")
		if err != nil {
			log.Fatalln(err)
		}
		if err := s.Start(addr); err != nil {
			log.Fatalln(err)
		}
	},
}
