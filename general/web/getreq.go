package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func getData() {

    resp, err := http.Get("https://www.tcmb.gov.tr/reeskontkur/202208/01082022-1300.xml")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}