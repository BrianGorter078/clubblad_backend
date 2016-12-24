package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

//Stores Clubblad information
//Number = Clubbladnummer
//URL = URL to Clubblad
type Clubblad struct {
	Number int
	URL    string
}

//Link to Clubblad
const CLUBBLAD_URL string = "http://www.kc-dordrecht.nl/wp-content/uploads/WB_2017_%s.pdf"

//Name of the file to write to
const FILE_NAME = "Clubbladen"

func looper(url string) {
	for clubbladNumber := 0; clubbladNumber < 20; clubbladNumber++ {
		httpGet(fmt.Sprintf(CLUBBLAD_URL, strconv.Itoa(clubbladNumber)), clubbladNumber)
	}
}

func httpGet(url string, clubbladNumber int) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(resp.Header)
	}

	if resp.StatusCode != 200 {
		return
	}

	persistResponse(url, clubbladNumber)
}

func persistResponse(responseBody string, clubbladNumber int) {
	file := fileLookUp()
	defer file.Close()

	jsonString := &Clubblad{
		Number: clubbladNumber,
		URL:    responseBody,
	}

	if !readFileContent(jsonString) {
		writeToFile(file, jsonString)
	}
}

func fileLookUp() *os.File {
	var file *os.File
	file, err := os.OpenFile(FILE_NAME, os.O_APPEND|os.O_RDWR, 0600)

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
	out, err := os.Create(FILE_NAME)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return out, nil
}

func readFileContent(jsonString *Clubblad) (exists bool) {
	fmt.Println("readfile")
	clubblad, err := json.Marshal(jsonString)
	if err != nil {
		fmt.Printf("error %v", err)
	}

	fileContent, readErr := ioutil.ReadFile(FILE_NAME)

	if readErr != nil {
		fmt.Println(readErr)
	}

	if bytes.Contains(fileContent, clubblad) {
		return true
	}

	return false
}

func writeToFile(file *os.File, jsonString *Clubblad) {
	err := json.NewEncoder(file).Encode(jsonString)
	if err != nil {
		fmt.Printf("Failed writing to disk! %v", err)
	}
}
