package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Minute)
	for {
		currentTime := time.Now().Local()
		fmt.Println(currentTime)
		looper(CLUBBLAD_URL)
		fmt.Println("Done")
		<-t.C
	}
}
