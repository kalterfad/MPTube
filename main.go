package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}

type VideoLink struct {
	Url string `json:"url"`
}

func main() {
	http.HandleFunc("/link", openMPV)
	log.Fatal(http.ListenAndServe(":9086", nil))
}

func openMPV(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var videoLink VideoLink
	err := decoder.Decode(&videoLink)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cmd := exec.Command("mpv", videoLink.Url)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	log.Printf("start video: '%v'", videoLink.Url)
	err = cmd.Run()
	log.Printf("end video: '%v'", videoLink.Url)
	if err != nil {
		log.Printf("cmd.Run err: '%v'", err)
	}
}
