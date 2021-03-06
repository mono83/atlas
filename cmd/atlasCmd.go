package cmd

import (
	"github.com/mono83/atlas/cmd/mysql"
	"github.com/spf13/cobra"
)

// AtlasCmd is main command for Atlas
var AtlasCmd = &cobra.Command{
	Use:   "atlas",
	Short: "Atlas database toolset",
}

func init() {
	AtlasCmd.AddCommand(
		mysql.Cmd,
	)
}
