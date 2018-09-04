package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	c "github.com/exiaohao/host-checker/pkg/client"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "host-checker",
	Short: "A host checker",
	Long:  "A host checker keeps unique hostname from istio gateway/virtualservice, prevent duplicated hostname caused access error",
	RunE: func(cmd *cobra.Command, args []string) error {
		stopCh := setupSignalHandler()
		watcher := new(c.Watcher)
		watcher.Init()
		watcher.Run(stopCh)
		return nil
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	fmt.Println("Execute called!")
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// fmt.Println("inited config")
}

// init
func init() {
	// fmt.Println("init")
	cobra.OnInitialize(initConfig)
}

// setupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler() (stopCh <-chan struct{}) {
	stop := make(chan struct{})
	sigs := make(chan os.Signal, 2)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		close(stop)
		<-sigs
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
