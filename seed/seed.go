package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Entry struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func main() {
	fmt.Println("Hello, World!")
	apiUrl := "https://memorize-it-api-4f185fe94534.herokuapp.com/entries/create"
	jsonFile, err := os.Open("seed_data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened seed_data.json")
	defer jsonFile.Close()
	var entries []Entry
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &entries)
	for _, entry := range entries {
		fmt.Println(entry.Name)
		fmt.Println(entry.Content)
		jsonData, err := json.Marshal(entry)
		if err != nil {
			fmt.Println(err)
		}

		req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}
