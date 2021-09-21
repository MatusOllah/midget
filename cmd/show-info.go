package cmd

import (
	"fmt"
	utils "midget/utils"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showInfoCmd = &cobra.Command{
	Use:   "show-info",
	Short: "Show info about a package.",
	Long:  `Download the package manifest and show the package info.`,

	//Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		b := color.New(color.FgCyan).SprintfFunc()

		if manifest != "" {
			pkg, err := utils.ReadManifest(manifest, viper.GetString("mods_folder"))
			utils.CheckError(err)

			if pkg.Installed {
				fmt.Printf("%v [installed]\n", b("%v - %v", pkg.Name, pkg.Version))
				fmt.Println()
				fmt.Printf("%v: %v\n", b("Full Name"), pkg.FullName)
				utils.PrintTriggers(pkg.Triggers)
			} else {
				color.Cyan("%v - %v", pkg.Name, pkg.Version)
				fmt.Println()
				fmt.Printf("%v: %v\n", b("Full Name"), pkg.FullName)
				utils.PrintTriggers(pkg.Triggers)
			}

		} else {
			pkg, err := utils.DownloadAndReadManifest(args[0], 0, viper.GetString("gh_repo"), viper.GetString("mods_folder"))
			utils.CheckError(err)

			if pkg.Installed {
				fmt.Printf("%v [installed]\n", b("%v - %v", pkg.Name, pkg.Version))
				fmt.Println()
				fmt.Printf("%v: %v\n", b("Full Name"), pkg.FullName)
				utils.PrintTriggers(pkg.Triggers)
			} else {
				color.Cyan("%v - %v", pkg.Name, pkg.Version)
				fmt.Println()
				fmt.Printf("%v: %v\n", b("Full Name"), pkg.FullName)
				utils.PrintTriggers(pkg.Triggers)
			}
		}
	},
}

func init() {
	showInfoCmd.Flags().StringVarP(&manifest, "manifest", "m", "", "manifest file")

	rootCmd.AddCommand(showInfoCmd)
}
