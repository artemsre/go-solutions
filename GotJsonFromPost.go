package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func getJsonFromPost(url string, jsonStr string) (string, error) {
	var jsonBt = []byte(jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBt))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func main() {
	jStr := `{"title":"Buy cheese and bread for breakfast."}`
	out, _ := getJsonFromPost("http://restapi3.apiary.io/notes", jStr)
	print(out)
}
