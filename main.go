package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var (
	Logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
		TimeFormat:      time.Kitchen,
	})
)

func fileExist(w http.ResponseWriter, r *http.Request, path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		http.Redirect(w,r, "https://rayna.tech", http.StatusMovedPermanently)
		return
	} 
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		args := len(r.URL.Path[1:])
		if args >= 1 {
			path := fmt.Sprintf("storage/%s", r.URL.Path[1:])
			fileExist(w,r, path)
			http.ServeFile(w, r, path)
		} else {
			http.Redirect(w, r, "https://rayna.tech", http.StatusMovedPermanently)
		}
	})

	Logger.Info(fmt.Sprintf("Server listening %s", ":8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		Logger.Fatal(err)
	}
}
