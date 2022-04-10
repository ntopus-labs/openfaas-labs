package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	fileBeingRead, err := ioutil.ReadFile("../example.csv")
	checkError("failed during read fileBeingRead", err)
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/function/csv-to-xlsx.openfaas-fn", bytes.NewReader(fileBeingRead))
	checkError("failed during create new post request", err)
	resp, err := http.DefaultClient.Do(req)
	checkError("failed during execute post request", err)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	checkError("failed during read post response body", err)
	createdFile, err := os.Create("./converted-example.xlsx")
	checkError("failed during create converted xlsx fileBeingRead", err)
	defer func() {
		err := createdFile.Close()
		checkError("failed during close fileBeingRead descriptor", err)
	}()
	_, err = createdFile.Write(bodyBytes)
	checkError("failed during write bytes into fileBeingRead", err)
	defer func() {
		err := resp.Body.Close()
		checkError("failed during close response body", err)
	}()
}

func checkError(message string, err error) {
	if err != nil {
		log.Printf("%s: %s", message, err.Error())
	}
}
