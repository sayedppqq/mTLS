package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	cert, err := ioutil.ReadFile("/home/appscodepc/go/src/github.com/sayedppqq/mTLS/certs/ca.crt")
	if err != nil {
		log.Fatalf("couldnt open ca.cart file")
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(cert)
	client := http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
	}

	resp, err := client.Get("https://localhost:9090")
	if err != nil {
		log.Fatalf("error while making request %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading response: %v", err)
	}
	fmt.Println(string(body))

}
