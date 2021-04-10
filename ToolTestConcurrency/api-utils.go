package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ApiUtil struct{}

func (api *ApiUtil) apiGetRequest(url string) string {
	timeout := time.Duration(time.Second * 10)
	client := http.Client{
		Timeout: timeout,
	}

	request, _ := http.NewRequest("GET", url, nil)
	//request.Header.Set("Authorization", "Bearer abcdefghijklmnopqrstuvwxyz0123456789")

	resp, err := client.Do(request)
	if checkError(err) {
		fmt.Println("Error request | client")
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	return string(jsonBytes)
}

func (api *ApiUtil) apiPutRequest(url string, postData map[string]interface{}) string {

	requestBody, _ := json.Marshal(postData)

	timeout := time.Duration(time.Second * 5)
	client := http.Client{
		Timeout: timeout,
	}

	request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	resp, err := client.Do(request)
	if checkError(err) {
		fmt.Println("Error request | client")
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	return string(jsonBytes)
}

func (api *ApiUtil) apiPutRequestX(url string, postData interface{}) string {

	requestBody, _ := json.Marshal(postData)

	timeout := time.Duration(time.Second * 5)
	client := http.Client{
		Timeout: timeout,
	}

	request, _ := http.NewRequest("PUT", url, bytes.NewBuffer(requestBody))

	resp, err := client.Do(request)
	if checkError(err) {
		fmt.Println("Error request | client")
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	return string(jsonBytes)
}

func (api *ApiUtil) apiPostRequest(url string, postData map[string]interface{}) string {

	requestBody, _ := json.Marshal(postData)

	timeout := time.Duration(time.Second * 5)
	client := http.Client{
		Timeout: timeout,
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	resp, err := client.Do(request)
	if checkError(err) {
		fmt.Println("Error request | client")
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	return string(jsonBytes)
}

func (api *ApiUtil) apiPostRequestX(url string, postData interface{}) string {

	requestBody, _ := json.Marshal(postData)

	timeout := time.Duration(time.Second * 5)
	client := http.Client{
		Timeout: timeout,
	}

	request, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))

	resp, err := client.Do(request)
	if checkError(err) {
		fmt.Println("Error request | client")
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	return string(jsonBytes)
}

func (api *ApiUtil) apiDeleteRequest(url string) string {
	timeout := time.Duration(time.Second * 5)
	client := http.Client{
		Timeout: timeout,
	}

	request, _ := http.NewRequest("DELETE", url, nil)

	resp, err := client.Do(request)
	if checkError(err) {
		fmt.Println("Error request | client")
	}
	defer resp.Body.Close()

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	return string(jsonBytes)
}
