package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tendermint/tendermint/version"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version infoooooooo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}
