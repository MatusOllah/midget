package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	utils "midget/utils"
)

func init() { rootCmd.AddCommand(removeCmd) }

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Uninstall a package.",
	Long:  `Uninstall a package and remove any traces of it.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Removing package %v...\n", args[0])

		err := os.RemoveAll(fmt.Sprintf("%v/%v", viper.GetString("mods_folder"), args[0]))
		utils.CheckError(err)

		color.Green("\a\n\u2705 successfully removed package %v!", args[0])

	},
}
