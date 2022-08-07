/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	Currency "github.com/masum0813/getcurrencyinfo/general/currency"
	"github.com/spf13/cobra"
)

// dolarCmd represents the dolar command
var dolarCmd = &cobra.Command{
	Use:   "dolar",
	Short: "Get requested date and time Dolar Price from TCMB",
	Long:  `Get requested date and time Dolar Price from TCMB`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("dolar called")
		Currency.GetCurrency("USD")
	},
}

var usdCmd = &cobra.Command{
	Use:   "usd",
	Short: "Get requested date and time value from TCMB",
	Long:  `Get requested date and time value from TCMB`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("dolar called")
		Currency.GetCurrency("USD")
	},
}

func init() {
	rootCmd.AddCommand(dolarCmd, usdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dolarCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dolarCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
