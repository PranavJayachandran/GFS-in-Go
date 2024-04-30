package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type dataFormat struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func hearBeat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heart beat call")
	fmt.Fprintf(w, "Heartbeat")
}
func attendance() {
	const serverUrl string = "http://localhost:8080/attendance"
	var myurl string = "http://localhost:" + os.Args[1]
	request := strings.NewReader(`
		{
			"id":"` + os.Args[1] + `",
			"url":"` + myurl + `"
		}
	`)
	resp, err := http.Post(serverUrl, "application/json", request)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
	defer resp.Body.Close()
}
func upload(w http.ResponseWriter, r *http.Request) {
	data := &dataFormat{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = os.MkdirAll(os.Args[1], 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	filePath := filepath.Join(os.Args[1], data.Name)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, strings.NewReader(data.Data))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("got Data:", data)
}
func main() {

	http.HandleFunc("/heartBeat", hearBeat)
	http.HandleFunc("/upload", upload)
	attendance()
	http.ListenAndServe(":"+os.Args[1], nil)
}
