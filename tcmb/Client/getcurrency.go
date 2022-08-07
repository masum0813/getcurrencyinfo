package GetCurrency

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	GetReq "github.com/masum0813/getcurrencyinfo/general/web"
	Doviz "github.com/masum0813/getcurrencyinfo/tcmb/Model"
)

var TCMBUrl = "https://www.tcmb.gov.tr/reeskontkur/%s/%s-%s.xml"

// the layout we want to use to parse our
// date string

func GetCurrency(date string, time string) Doviz.Doviz {
	// getreq.GetData()

	yearMonth := GetYearMonth(date)
	dayMonthYear := GetDayMonthYear(date)

	requestUrL := fmt.Sprintf(TCMBUrl, yearMonth, dayMonthYear, time)

	body := GetReq.GetData(requestUrL)

	doviz := Doviz.Doviz{}
	err := xml.Unmarshal([]byte(body), &doviz)
	// Wraps the response to Response struct
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))

	// fmt.Printf("%s\n", doviz.BaslikBilgi.Yayimlayan)

	// for i, item := range doviz.DovizKurListe.Kurs {
	// 	fmt.Printf("%d: %s\n", i, item.DovizCinsi)
	// }

	return doviz
}

func GetYearMonth(date string) string {

	// we parse our date string to our layout
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	month := "00"
	if int(t.Month()) < 10 {
		month = "0" + fmt.Sprintf("%d", int(t.Month()))
	}
	// fmt.Printf("%d%s\n", t.Year(), month)
	return fmt.Sprintf("%d%s", t.Year(), month)

}

func GetDayMonthYear(date string) string {
	// we parse our date string to our layout
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		log.Fatal(err)
	}
	month := "00"
	if int(t.Month()) < 10 {
		month = "0" + fmt.Sprintf("%d", int(t.Month()))
	}

	day := "00"
	if int(t.Day()) < 10 {
		day = "0" + fmt.Sprintf("%d", int(t.Day()))
	}

	// fmt.Printf("%s%s%d\n", day, month, t.Year())
	return fmt.Sprintf("%s%s%d", day, month, t.Year())
}
