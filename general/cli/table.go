package table

import (
	"fmt"
	"os"

	DTProcess "github.com/masum0813/getcurrencyinfo/general/datetime"
	Doviz "github.com/masum0813/getcurrencyinfo/tcmb/Model"
	"github.com/olekukonko/tablewriter"
)

func GenerateTable(doviz Doviz.Doviz, currency string) {

	caption := "\n"
	currencyDetail := ""
	caption += fmt.Sprintf("%s\n", doviz.BaslikBilgi.Yayimlayan)
	// caption += fmt.Sprintf("Tel: %s\n", doviz.BaslikBilgi.Tel)
	// caption += fmt.Sprintf("Faks: %s\n", doviz.BaslikBilgi.Faks)
	// caption += fmt.Sprintf("EPosta: %s\n", doviz.BaslikBilgi.Eposta)
	caption += fmt.Sprintf("Kur Tarih: %s\n", DTProcess.GetCustomRequestDate(doviz.BaslikBilgi.ZamanEtiketi))
	caption += fmt.Sprintf("Kur Saat: %s\n", DTProcess.GetCustomRequestTime(doviz.BaslikBilgi.ZamanEtiketi))
	fmt.Println(caption)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Doviz Kodu", "Doviz Aciklama", "Tutar"})

	for _, v := range doviz.DovizKurListe.Kurs {
		switch v.DovizCinsi {
		case "USD":
			currencyDetail = "Amerikan Doları"
		case "EUR":
			currencyDetail = "Euro"
		case "GBP":
			currencyDetail = "İngiliz Sterlini"
		case "CHF":
			currencyDetail = "İsviçre Frangı"
		case "XAU":
			currencyDetail = "Altın 995/1000"
		case "XAS":
			currencyDetail = "Altın 1000/1000"
		default:
			currencyDetail = "Not Found"

		}
		if v.DovizCinsi == currency || currency == "All" {
			table.Append([]string{v.DovizCinsi, currencyDetail, fmt.Sprintf("%s₺", v.Alis)})
		}
	}

	table.Render() // Send output
}
