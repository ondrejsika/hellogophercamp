package hello

import (
	"fmt"

	"github.com/ondrejsika/hellogophercamp/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "hello",
	Short:   "Hello, GopherCamp!",
	Aliases: []string{"h"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Hello Gophercamp\n")
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
