package cmd

import (
	"fmt"
	utils "midget/utils"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a package.",
	Long:  `Download and upgrade a package.`,

	Run: func(cmd *cobra.Command, args []string) {
		pathc := make(chan string)
		errc := make(chan error)

		pathc <- fmt.Sprintf("%v/%v", viper.GetString("mods_folder"), args[0])
		close(pathc)

		go func() {
			err := os.RemoveAll(<-pathc)
			if err != nil {
				errc <- err
			}
		}()

		for {
			err, open := <-errc

			fmt.Printf("\rRemoving package %v... |", args[0])
			time.Sleep(100 * time.Millisecond)

			fmt.Printf("\rRemoving package %v... /", args[0])
			time.Sleep(100 * time.Millisecond)

			fmt.Printf("\rRemoving package %v... -", args[0])
			time.Sleep(100 * time.Millisecond)

			fmt.Printf("\rRemoving package %v... \\", args[0])
			time.Sleep(100 * time.Millisecond)

			if !open {
				utils.CheckError(err)
				break
			}
		}

		pkg, merr := utils.DownloadAndReadManifest(args[0], 0, viper.GetString("gh_repo"), viper.GetString("mods_folder"))
		utils.CheckError(merr)

		utils.PrintTriggers(pkg.Triggers)

		err := pkg.InstallPackage(viper.GetString("mods_folder"))
		utils.CheckError(err)

		color.Green("\u2705 successfully upgraded package %v to %v!", pkg.Name, pkg.Version)
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
