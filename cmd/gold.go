/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	Currency "github.com/masum0813/getcurrencyinfo/general/currency"
	"github.com/spf13/cobra"
)

// goldCmd represents the gold command
var goldCmd = &cobra.Command{
	Use:   "gold",
	Short: "Get requested date and time Gold(995, 1000) Price from TCMB",
	Long:  `Get requested date and time Gold(995, 1000) Price from TCMB`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("gold called")
		Currency.GetCurrency("GOLD")
	},
}

func init() {
	rootCmd.AddCommand(goldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
