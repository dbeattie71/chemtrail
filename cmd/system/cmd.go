package system

import (
	"fmt"
	"os"

	"github.com/jrasell/chemtrail/cmd/system/health"
	"github.com/sean-/sysexits"
	"github.com/spf13/cobra"
)

func RegisterCommand(rootCmd *cobra.Command) error {
	cmd := &cobra.Command{
		Use:   "system",
		Short: "Retrieve system information about a Chemtrail server",
		Run: func(cmd *cobra.Command, args []string) {
			runSystem(cmd, args)
		},
	}
	rootCmd.AddCommand(cmd)

	if err := registerCommands(cmd); err != nil {
		fmt.Println("Error registering commands:", err)
		os.Exit(sysexits.Software)
	}
	return nil
}

func runSystem(cmd *cobra.Command, _ []string) {
	_ = cmd.Usage()
}

func registerCommands(cmd *cobra.Command) error {
	return health.RegisterCommand(cmd)
}
