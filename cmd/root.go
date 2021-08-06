package cmd

import (
	"github.com/endigma/garma/ui"
	"github.com/spf13/cobra"
)

// cmd_root represents the base command when called without any subcommands
var cmd_root = &cobra.Command{
	Use:     "gaa [command] [args...] [flags]",
	Version: "v0.1.6",
	Short:   "gaa is a gtk/cli launcher for Arma 3 on linux",
	Run: func(cmd *cobra.Command, args []string) {
		ui.Run()
	},
}

func Execute() {
	cobra.CheckErr(cmd_root.Execute())
}
