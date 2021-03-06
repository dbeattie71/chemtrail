package delete

import (
	"fmt"
	"os"

	"github.com/jrasell/chemtrail/pkg/api"
	"github.com/jrasell/chemtrail/pkg/config/client"
	"github.com/sean-/sysexits"
	"github.com/spf13/cobra"
)

func RegisterCommand(rootCmd *cobra.Command) error {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes a scaling policy",
		Run: func(cmd *cobra.Command, args []string) {
			runDelete(cmd, args)
		},
	}
	rootCmd.AddCommand(cmd)

	return nil
}

func runDelete(_ *cobra.Command, args []string) {
	switch {
	case len(args) < 1:
		fmt.Println("Not enough arguments, expected 1 args got", len(args))
		os.Exit(sysexits.Usage)
	case len(args) > 1:
		fmt.Println("Too many arguments, expected 1 args got", len(args))
		os.Exit(sysexits.Usage)
	}

	clientConfig := client.GetConfig()
	mergedConfig := api.DefaultConfig(&clientConfig)

	chemtrailClient, err := api.NewClient(mergedConfig)
	if err != nil {
		fmt.Println("Error setting up Chemtrail client:", err)
		os.Exit(sysexits.Software)
	}

	if err := chemtrailClient.Policy().Delete(args[0]); err != nil {
		fmt.Println("Error deleting class scaling policy:", err)
		os.Exit(sysexits.Software)
	}
	fmt.Println("Successfully deleted client scaling policy")
}
