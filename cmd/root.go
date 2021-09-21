package cmd

import (
	"fmt"
	utils "midget/utils"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var manifest string

var rootCmd = &cobra.Command{
	Use:   "midget",
	Short: "midget is a package manager for FNF mods",
	Long:  `midget is a package manager for Friday Night Funkin\' mods`,
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.midgetrc)")

	rootCmd.SilenceErrors = true
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		utils.CheckError(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".midgetrc")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		b := color.New(color.FgCyan).SprintFunc()

		fmt.Printf("%s: %s\n", b("Using config file"), viper.ConfigFileUsed())
		fmt.Println()
	}
}

func Execute() {
	utils.CheckError(rootCmd.Execute())
}
