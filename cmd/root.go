package cmd

import (
	"fmt"
	"os"

	c "github.com/exiaohao/host-chekcer/pkg/client"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "host-checker",
	Short: "A host checker",
	Long:  "A host checker keeps unique hostname from istio gateway/virtualservice, prevent duplicated hostname caused access error",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RootCmd: RUN")
		c.Run()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("Execute called!")
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	fmt.Println("inited config")
}

// init
func init() {
	fmt.Println("init")
	cobra.OnInitialize(initConfig)
}
