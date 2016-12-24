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
		looper(CLUBBLADURL)
		<-t.C
	}
}
