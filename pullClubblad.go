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

const CLUBBLADURL = "http://www.kc-dordrecht.nl/wp-content/uploads/WB_2017_%s.pdf"

func looper(url string) {
	for index := 0; index < 20; index++ {
		httpGet(fmt.Sprintf(CLUBBLADURL, strconv.Itoa(index)), index)
	}
}

func httpGet(url string, index int) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(resp.Header)
	}

	if resp.StatusCode != 200 {
		return
	}

	var respCode = strconv.Itoa(resp.StatusCode)
	fmt.Println(respCode)
	persistResponse(url, index)
}

func persistResponse(responseBody string, index int) {
	out := checkFile()
	defer out.Close()

	jsonString := &Clubblad{
		Number: index,
		URL:    responseBody,
	}

	if err := json.NewEncoder(out).Encode(jsonString); err != nil {
		fmt.Printf("Failed writing to disk! %v", err)
	}
}

func checkFile() *os.File {
	var file *os.File
	file, err := os.OpenFile("Clubbladen", os.O_APPEND|os.O_WRONLY, 0600)

	if os.IsNotExist(err) {
		file, err = createFile()
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}
	return file
}

func createFile() (*os.File, error) {
	out, err := os.Create("Clubbladen")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return out, nil
}
