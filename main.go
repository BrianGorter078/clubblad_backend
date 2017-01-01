package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Stores Clubblad information
//Number = Clubbladnummer
//URL = URL to Clubblad
type Clubblad struct {
	Number int
	URL    string
}

var availableClubbladen = []Clubblad{}

//Link to Clubblad
const CLUBBLAD_URL string = "http://www.kc-dordrecht.nl/wp-content/uploads/WB_2017_%s.pdf"

func main() {

	go timer()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func timer() {
	t := time.NewTicker(time.Minute)
	for {
		currentTime := time.Now().Local()
		fmt.Println(currentTime)
		looper(CLUBBLAD_URL)
		fmt.Println("Done")
		<-t.C
	}
}

func looper(url string) {
	availableClubbladen = []Clubblad{}

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

	} else {
		availableClubbladen = append(availableClubbladen, Clubblad{clubbladNumber, url})
		fmt.Println(availableClubbladen)
	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(availableClubbladen)
}
