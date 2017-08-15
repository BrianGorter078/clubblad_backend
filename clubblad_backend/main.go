package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"os"
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
const CLUBBLAD_URL string = "http://www.kc-dordrecht.nl/wp-content/uploads/WB_2018_%s.pdf"

func main() {
	//Setting the port from a environment variable to listen on on heroku
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//Looping over all posible url's to get all the availableClubbladen
	go timer()

	http.HandleFunc("/kcd", kcd)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
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

	for clubbladNumber := 30; clubbladNumber > 0; clubbladNumber-- {
		var number = clubbladNumber
		if clubbladNumber < 10 {

			httpGet(fmt.Sprintf(CLUBBLAD_URL, leftPad(strconv.Itoa(clubbladNumber), "0", 1)), clubbladNumber)
		} else {
			httpGet(fmt.Sprintf(CLUBBLAD_URL, strconv.Itoa(number)), clubbladNumber)
		}

	}
}

func leftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}

func httpGet(url string, clubbladNumber int) {
	resp, err := http.Get(url)

	fmt.Println(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Println(resp.Header)
		return
	}

	loadClubbladen = append(loadClubbladen, Clubblad{
		Number: clubbladNumber,
		URL:    url,
	})

	fmt.Println(loadClubbladen)

}

func kcd(writer http.ResponseWriter, r *http.Request) {
	json.NewEncoder(writer).Encode(availableClubbladen)
}
