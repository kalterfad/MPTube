package main

import (
	"encoding/json"
	"fmt"
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
	err := notification("Start MPTube")
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/link", openMPV)
	log.Fatal(http.ListenAndServe(":9086", nil))
}

func notification(message string) error {
	cmd := exec.Command(
		"notify-send",
		"-i", "/usr/share/icons/MPTube/icon.svg",
		"MPTube", message)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("notfication err:%v", err)
	}
	return nil
}

func openMPV(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var videoLink VideoLink
	err := decoder.Decode(&videoLink)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = notification("Start video")
	cmd := exec.Command("mpv", videoLink.Url)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Run()

	if err != nil {
		_ = notification("Open error: \n" + err.Error())
		return
	}
}
