package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"os"

	"github.com/gorilla/mux"
)

//Stores Clubblad information
//Number = Clubbladnummer
//URL = URL to Clubblad
type Clubblad struct {
	Title  string `json:"title,omitempty"`
	Number int    `json:"number"`
	URL    string `json:"url"`
}

var availableClubbladen = []Clubblad{}
var loadClubbladen = []Clubblad{}

//Link to Clubblad
const CLUBBLAD_URL string = "http://www.kc-dordrecht.nl/wp-content/uploads/WB_2017_%s.pdf"

func main() {

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	go timer()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

func timer() {
	t := time.NewTicker(time.Minute)
	for {
		currentTime := time.Now().Local()
		fmt.Println(currentTime)
		looper(CLUBBLAD_URL)
		availableClubbladen = loadClubbladen
		fmt.Println("Done")
		<-t.C
	}
}

func looper(url string) {
	loadClubbladen = []Clubblad{}

	for clubbladNumber := 20; clubbladNumber > 0; clubbladNumber-- {
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
		loadClubbladen = append(loadClubbladen, Clubblad{
			Number: clubbladNumber,
			URL:    url,
		})
		fmt.Println(loadClubbladen)
	}

}

func Index(writer http.ResponseWriter, r *http.Request) {
	json.NewEncoder(writer).Encode(availableClubbladen)
}
