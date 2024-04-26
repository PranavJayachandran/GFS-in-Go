package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	f, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	for {
		chunk, err := io.ReadAll(io.LimitReader(file, byteSize))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(chunk) == 0 { //Else this goes on tho be a infinite loop
			break
		}
		sendToChunkServer(string(chunk))
	}

	json.NewEncoder(w).Encode(message{Msg: "File uploaded succesfukly"})
}
func sendToChunkServer(chunk string) {
	request := strings.NewReader(`
		{
			"name":"` + generateName() + `",
			"data":"` + chunk + `"
		}
	`)
	serverUrl := chunkServerMap[serverNames[rand.Intn(len(serverNames))]] + "/upload"
	resp, err := http.Post(serverUrl, "application/json", request)
	if err != nil {
		panic(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content), "Sent to ", serverUrl)
	defer resp.Body.Close()
}
