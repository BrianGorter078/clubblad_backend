package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Clubblad struct {
	Number int
	URL    string
}

const CLUBBLADURL = "http://www.kc-dordrecht.nl/wp-content/uploads/WB_2017_"

func looper(url string) {
	for index := 0; index < 20; index++ {
		var indexString = strconv.Itoa(index)
		httpGet(url+indexString+".pdf", index)
	}
}

func httpGet(url string, index int) {
	resp, err := http.Get(url)

	if resp.StatusCode != 200 {
	} else {
		var respCode = strconv.Itoa(resp.StatusCode)
		fmt.Println(respCode)
		persistResponse(url, index)
	}

	if err != nil {
		fmt.Println(resp.Header)
	}
}

func persistResponse(responseBody string, index int) {
	out := checkFile()
	jsonString := &Clubblad{
		Number: index,
		URL:    responseBody}
	jsonResponse, _ := json.Marshal(jsonString)
	fmt.Println(string(jsonResponse))
	out.WriteString(string(jsonResponse) + "\n")

}

func checkFile() *os.File {
	out, err := os.OpenFile("Clubbladen", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		createdFile, err := createFile()
		if err != nil {
			fmt.Println(err)
		} else {
			return createdFile
		}
	}
	return out
}

func createFile() (*os.File, error) {
	out, err := os.Create("Clubbladen")

	if err != nil {
		fmt.Println(err)
		out.Close()
		return nil, err
	}

	return out, nil
}
