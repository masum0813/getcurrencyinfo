/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	Currency "github.com/masum0813/getcurrencyinfo/general/currency"
	"github.com/spf13/cobra"
)

// euroCmd represents the euro command
var euroCmd = &cobra.Command{
	Use:   "euro",
	Short: "Get requested date and time Euro Price from TCMB",
	Long:  `Get requested date and time Euro Price from TCMB`,

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("euro called")
		Currency.GetCurrency("EUR")
	},
}

var eurCmd = &cobra.Command{
	Use:   "eur",
	Short: "Get requested date and time Euro Price from TCMB",
	Long:  `Get requested date and time Euro Price from TCMB`,

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("euro called")
		Currency.GetCurrency("EUR")
	},
}

func init() {
	rootCmd.AddCommand(euroCmd, eurCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// euroCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// euroCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
