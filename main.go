package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/thedevsaddam/gojsonq"
)

var (
	url       = "https://xxxxxxx"
	auth      = "Bearer xxxxxxxx"
	action    = "GET"
	transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client    = &http.Client{Transport: transport}
)

// appel de l'api
func callApi(api_action, api_url, api_auth string) (string, error) {

	// create request
	req, err := http.NewRequest(api_action, api_url, nil)
	if err != nil {
		return "", err
	}

	// set auth
	req.Header.Add("Authorization", api_auth)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// get response
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func main() {

	str, err := callApi(action, url, auth)
	if err != nil {
		fmt.Println(err)
		return
	}

	// SAMPLE
	jq := gojsonq.New().JSONString(str).From("FOO").WhereEqual("BAR", true).Find("[0].USER.NAME")

	// print
	fmt.Println(jq)
}
