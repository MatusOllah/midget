package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	utils "midget/utils"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update package list.",
	Long:  `Download and update package list.`,

	Run: func(cmd *cobra.Command, args []string) {
		repo := viper.GetString("gh_repo")
		tempDir := os.TempDir()

		url := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/packages.csv", repo)
		path := fmt.Sprintf("%s/midget/packages.csv", tempDir)

		fmt.Printf("Downloading package list from GitHub repo %s...\n", repo)
		cobra.CheckErr(utils.DownloadFile(url, path))
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
