package root

import (
	"fmt"

	"github.com/ondrejsika/hellogophercamp/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "hellogophercamp",
	Short: "hellogophercamp, " + version.Version,
	Run: func(c *cobra.Command, args []string) {
		fmt.Printf("Hello Gophercamp\n")
	},
}
