package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	bootDelay, _ := strconv.Atoi(os.Getenv("BOOT_DELAY"))
	time.Sleep(time.Duration(bootDelay) * time.Second)

	respDelay, _ := strconv.Atoi(os.Getenv("RESPONSE_DELAY"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if respDelay > 0 {
			time.Sleep(time.Duration(respDelay) * time.Second)
		}
		w.Write([]byte("OK\n"))
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
