package doviz

import (
	"encoding/xml"
)

type Doviz struct {
	XMLName       xml.Name      `xml:"tcmbVeri"`
	BaslikBilgi   BaslikBilgi   `xml:"baslik_bilgi"`
	DovizKurListe DovizKurListe `xml:"doviz_kur_liste"`
	Aciklama      Aciklama      `xml:"aciklama"`
}

type BaslikBilgi struct {
	XMLName      xml.Name `xml:"baslik_bilgi"`
	Kod          string   `xml:"kod"`
	VeriTipi     string   `xml:"veri_tipi"`
	VeriTanim    string   `xml:"veri_tanim"`
	Yayimlayan   string   `xml:"yayimlayan"`
	Tel          string   `xml:"tel"`
	Faks         string   `xml:"faks"`
	Eposta       string   `xml:"eposta"`
	ZamanEtiketi string   `xml:"zaman_etiketi"`
}

type DovizKurListe struct {
	XMLName xml.Name `xml:"doviz_kur_liste"`
	Kurs    []Kur    `xml:"kur"`
}

type Kur struct {
	XMLName          xml.Name `xml:"kur"`
	DovizCinsiTabani string   `xml:"doviz_cinsi_tabani"`
	DovizCinsi       string   `xml:"doviz_cinsi"`
	Birim            int      `xml:"birim"`
	Alis             string   `xml:"alis"`
	SiraNo           int      `xml:"sira_no"`
}

type Aciklama struct {
	XMLName  xml.Name `xml:"aciklama"`
	Aciklama string   `xml:"aciklama"`
}
