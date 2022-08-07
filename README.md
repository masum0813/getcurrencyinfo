# getcurrencyinfo

Get Currency Info from TCMB

## Usage

```bash
❯ getcurrencyinfo --help
Get Currency Info From TCMB

Usage:
  getcurrencyinfo [flags]
  getcurrencyinfo [command]

Available Commands:
  chf         Get requested date and time Frank Price from TCMB
  completion  Generate the autocompletion script for the specified shell
  dolar       Get requested date and time Dolar Price from TCMB
  eur         Get requested date and time Euro Price from TCMB
  euro        Get requested date and time Euro Price from TCMB
  frank       Get requested date and time Frank Price from TCMB
  gbp         Get requested date and time Sterlin Price from TCMB
  gold        Get requested date and time Gold(995, 1000) Price from TCMB
  help        Help about any command
  sterlin     Get requested date and time Sterlin Price from TCMB
  usd         Get requested date and time value from TCMB
  version     Show getcurrencyinfo version information

Flags:
  -h, --help     help for getcurrencyinfo
  -t, --toggle   Help message for toggle

Use "getcurrencyinfo [command] --help" for more information about a command.
```

## Examples

Get all currency prices from TCMB:

```bash
❯ getcurrencyinfo

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+------------------+----------+
| DOVIZ KODU |  DOVIZ ACIKLAMA  |  TUTAR   |
+------------+------------------+----------+
| EUR        | Euro             | 18,3879₺ |
| USD        | Amerikan Doları  | 17,9727₺ |
| GBP        | İngiliz Sterlini | 21,8189₺ |
| CHF        | İsviçre Frangı   | 18,8117₺ |
| XAU        | Altın 995/1000   | 1032,44₺ |
| XAS        | Altın 1000/1000  | 1037,63₺ |
+------------+------------------+----------+
```

Get Dolar Price from TCMB:

```bash
❯ getcurrencyinfo dolar

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+-----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA  |  TUTAR   |
+------------+-----------------+----------+
| USD        | Amerikan Doları | 17,9727₺ |
+------------+-----------------+----------+
```

```bash
❯ getcurrencyinfo usd

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+-----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA  |  TUTAR   |
+------------+-----------------+----------+
| USD        | Amerikan Doları | 17,9727₺ |
+------------+-----------------+----------+
```

Get Euro Price from TCMB:

```bash
❯ getcurrencyinfo euro

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA |  TUTAR   |
+------------+----------------+----------+
| EUR        | Euro           | 18,3879₺ |
+------------+----------------+----------+
```

```bash
❯ getcurrencyinfo eur

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA |  TUTAR   |
+------------+----------------+----------+
| EUR        | Euro           | 18,3879₺ |
+------------+----------------+----------+
```

Get Frank Price from TCMB:

```bash
❯ getcurrencyinfo frank

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA |  TUTAR   |
+------------+----------------+----------+
| CHF        | İsviçre Frangı | 18,8117₺ |
+------------+----------------+----------+
```

```bash
❯ getcurrencyinfo chf

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA |  TUTAR   |
+------------+----------------+----------+
| CHF        | İsviçre Frangı | 18,8117₺ |
+------------+----------------+----------+
```

Get Sterlin Price from TCMB:

```bash
❯ getcurrencyinfo sterlin

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+------------------+----------+
| DOVIZ KODU |  DOVIZ ACIKLAMA  |  TUTAR   |
+------------+------------------+----------+
| GBP        | İngiliz Sterlini | 21,8189₺ |
+------------+------------------+----------+
```

```bash
❯ getcurrencyinfo gbp

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+------------------+----------+
| DOVIZ KODU |  DOVIZ ACIKLAMA  |  TUTAR   |
+------------+------------------+----------+
| GBP        | İngiliz Sterlini | 21,8189₺ |
+------------+------------------+----------+
```

Get Gold Price from TCMB:

```bash
❯ getcurrencyinfo gold

TCMB Piyasalar Genel Müdürlüğü - Döviz Piyasaları Müdürlüğü
Kur Tarih: 05.08.2022
Kur Saat: 15:01

+------------+-----------------+----------+
| DOVIZ KODU | DOVIZ ACIKLAMA  |  TUTAR   |
+------------+-----------------+----------+
| XAU        | Altın 995/1000  | 1032,44₺ |
| XAS        | Altın 1000/1000 | 1037,63₺ |
+------------+-----------------+----------+
```

## Current Currency TCMB

TCMB currency and gold price list urls:

<https://www.tcmb.gov.tr/reeskontkur/202206/01062022-1300.xml>

<https://www.tcmb.gov.tr/reeskontkur/202208/01082022-1300.xml>

### Create Table

Create bash table :

<https://github.com/olekukonko/tablewriter>
