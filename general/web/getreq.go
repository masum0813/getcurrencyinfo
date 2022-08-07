package getreq

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetData(requestUrL string) string {

	response, err := http.Get(requestUrL)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
