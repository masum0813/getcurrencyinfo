/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	Currency "github.com/masum0813/getcurrencyinfo/general/currency"
	"github.com/spf13/cobra"
)

// sterlinCmd represents the sterlin command
var sterlinCmd = &cobra.Command{
	Use:   "sterlin",
	Short: "Get requested date and time Sterlin Price from TCMB",
	Long:  `Get requested date and time Sterlin Price from TCMB`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("sterlin called")
		Currency.GetCurrency("GBP")
	},
}

var gbpCmd = &cobra.Command{
	Use:   "gbp",
	Short: "Get requested date and time Sterlin Price from TCMB",
	Long:  `Get requested date and time Sterlin Price from TCMB`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("sterlin called")
		Currency.GetCurrency("GBP")
	},
}

func init() {
	rootCmd.AddCommand(sterlinCmd, gbpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sterlinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sterlinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
