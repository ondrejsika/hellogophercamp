package cmd

import (
	_ "github.com/ondrejsika/hellogophercamp/cmd/hello"
	"github.com/ondrejsika/hellogophercamp/cmd/root"
	_ "github.com/ondrejsika/hellogophercamp/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
