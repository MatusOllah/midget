package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	utils "midget/utils"
)

var version float64

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a package.",
	Long:  `Download and install a package.`,

	//Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		if manifest != "" {
			pkg, merr := utils.ReadManifest(manifest, viper.GetString("mods_folder"))
			utils.CheckError(merr)

			if pkg.Installed {
				utils.LogError(fmt.Sprintf("package %s already installed", pkg.Name))
			} else {
				utils.PrintTriggers(pkg.Triggers)

				err := pkg.InstallPackage(viper.GetString("mods_folder"))
				utils.CheckError(err)

				color.Green("\u2705 successfully installed package %v!", pkg.Name)
			}
		} else {
			pkg, merr := utils.DownloadAndReadManifest(args[0], version, viper.GetString("gh_repo"), viper.GetString("mods_folder"))
			utils.CheckError(merr)

			if pkg.Installed {
				utils.LogError(fmt.Sprintf("package %s already installed", pkg.Name))
			} else {
				utils.PrintTriggers(pkg.Triggers)

				err := pkg.InstallPackage(viper.GetString("mods_folder"))
				utils.CheckError(err)

				color.Green("\a\n\u2705 successfully installed package %v!", pkg.Name)
			}
		}
	},
}

func init() {
	installCmd.Flags().Float64VarP(&version, "version", "v", 0, "package version to install")
	installCmd.Flags().StringVarP(&manifest, "manifest", "m", "", "manifest file")

	rootCmd.AddCommand(installCmd)
}
