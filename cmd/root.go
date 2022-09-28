package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile = "ismyteamawake"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "is-my-team-awake",
	Short: "Tools to check the timezones of a remote team",
	Long: `Tools to check the timezones of a remote team, for example:
	
		is-my-team-awake check "Jonh Doe"
	`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Is my team awake cli -- HEAD")
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.ismyteamawake.yaml)")
	_ = viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))

	// commands
	rootCmd.AddCommand(check)
	rootCmd.AddCommand(list)
}

func initConfig() {
	viper.SetConfigName("default") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file: ", viper.ConfigFileUsed())
	}
}
