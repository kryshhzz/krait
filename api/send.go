package api

import (
	"log"
	"net/http"
	"bytes"
)

func SendReq(link string, from string) (*http.Response, error) { 

	log.Println("sending .... ")

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")

	if from == "PAYTM" {
		req.Header.Set("Referer", "https://tickets.paytm.com/trains")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-")
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	
	} else if from == "CTKT" {
		req.Header.Set("Accept", "*/*")
		req.Header.Set("Referer", "https://www.confirmtkt.com/") 
		req.Header.Set("Origin", "https://www.confirmtkt.com/") 
		req.Header.Set("Host", "securedapi.confirmtkt.com/") 
		req.Header.Set("Content-Type", "application/json")

	} else {
		// req.Header.Set("Host","trainticketapi.railyatri.in")
		req.Header.Set("Accept","*/*")
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	return resp, err
}


func SendPostReq(link string, from string, jsonData string) (*http.Response, error) {  

	log.Println("sending .... ")  

	req, err := http.NewRequest("POST", link, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0")

	if from == "IBIBO" {
		req.Header.Set("Host","rails-cbe.goibibo.com")
		req.Header.Set("Accept","*/*")
	} else {
		// req.Header.Set("Host","rails-cbe.goibibo.com")
		req.Header.Set("Accept","*/*")
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	return resp, err

}