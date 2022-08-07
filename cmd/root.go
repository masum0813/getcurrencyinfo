/*
Copyright © 2022 Murat Oğuz muratoguz@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"time"

	Cli "github.com/masum0813/getcurrencyinfo/general/cli"
	DateTimeProcess "github.com/masum0813/getcurrencyinfo/general/datetime"
	Client "github.com/masum0813/getcurrencyinfo/tcmb/Client"
	"github.com/spf13/cobra"
)

var (
	// VERSION is set during build
	VERSION string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "getcurrencyinfo",
	Short: "Get Currency Info From TCMB",
	Long:  `Get Currency Info From TCMB`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		currency := "All"
		varTime := time.Now()
		date := DateTimeProcess.GetDate(varTime.Format(time.RFC3339))
		retValTime := DateTimeProcess.GetTime(varTime.Format(time.RFC3339))
		doviz := Client.GetCurrency(date, retValTime)
		Cli.GenerateTable(doviz, currency)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	VERSION = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.getcurrencyinfo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
