package generatecurrency

import (
	"time"

	Cli "github.com/masum0813/getcurrencyinfo/general/cli"
	DateTimeProcess "github.com/masum0813/getcurrencyinfo/general/datetime"
	Client "github.com/masum0813/getcurrencyinfo/tcmb/Client"
)

func GetCurrency(currency string) {

	varTime := time.Now()
	date := DateTimeProcess.GetDate(varTime.Format(time.RFC3339))
	retValTime := DateTimeProcess.GetTime(varTime.Format(time.RFC3339))
	doviz := Client.GetCurrency(date, retValTime)
	Cli.GenerateTable(doviz, currency)
}
