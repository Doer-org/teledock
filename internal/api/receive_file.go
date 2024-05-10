package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Port string `json:"port"`
	ID   string `json:"id"`
}

func GetServerInfo(host string, id string) (string, error) {
	if host == "" {
		host = BackendURL
	}
	fullURL := fmt.Sprintf("%s/info/%s", host, id)
	fmt.Println(fullURL)
	response, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var responseBody bytes.Buffer
	_, err = io.Copy(&responseBody, response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(responseBody.String())

	var resp Response
	err = json.Unmarshal(responseBody.Bytes(), &resp)
	if err != nil {
		return "", err
	}
	fmt.Printf("%s/info/%s\n", FrontURL, id)
	return resp.Port, nil
}

func ReceiveTarGzFromServer(host string, id string) (string, error) {
	if host == "" {
		host = BackendURL
	}
	fullURL := fmt.Sprintf("%s/%s", host, id)
	response, err := http.Get(fullURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	infoURL := fmt.Sprintf("%s/info/%s", host, id)
	infoResponse, err := http.Get(infoURL)
	if err != nil {
		return "", err
	}
	defer infoResponse.Body.Close()

	port, err := GetServerInfo(host, id)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return port, nil
}
