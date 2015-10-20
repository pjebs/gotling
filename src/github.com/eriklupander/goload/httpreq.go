package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"time"
)


// Accepts a Httpaction and a one-way channel to write the results to.
func DoHttpRequest(httpAction HttpReq, resultsChannel chan StatFrame, simulationStart time.Time) {
	req, err := http.NewRequest(httpAction.Method, httpAction.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Accept", httpAction.Accept)


	client := &http.Client{}
	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	responseBody, err := ioutil.ReadAll(resp.Body)
	contentLength := len(responseBody)
	status := resp.StatusCode
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %d Content-Length: %d Latency %dms %s\n", status, contentLength, elapsed/1000000, httpAction.Url)
	statFrame := StatFrame{
		time.Since(simulationStart).Nanoseconds(),
		5,
	}
	resultsChannel <- statFrame
}