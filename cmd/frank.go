/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	Currency "github.com/masum0813/getcurrencyinfo/general/currency"
	"github.com/spf13/cobra"
)

// frankCmd represents the frank command
var frankCmd = &cobra.Command{
	Use:   "frank",
	Short: "Get requested date and time Frank Price from TCMB",
	Long:  `Get requested date and time Frank Price from TCMB`,

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("frank called")
		Currency.GetCurrency("CHF")
	},
}

var chfCmd = &cobra.Command{
	Use:   "chf",
	Short: "Get requested date and time Frank Price from TCMB",
	Long:  `Get requested date and time Frank Price from TCMB`,

	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("frank called")
		Currency.GetCurrency("CHF")
	},
}

func init() {
	rootCmd.AddCommand(frankCmd, chfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// frankCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// frankCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
