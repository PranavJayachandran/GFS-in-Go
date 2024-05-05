package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
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

	var folderPath []string
	folderPathStr := r.FormValue("folderPath")
	err = json.Unmarshal([]byte(folderPathStr), &folderPath)
	if err != nil {
		http.Error(w, "Failed to read file content", http.StatusInternalServerError)
		return
	}

	addFile(createFileFolderType{Path: folderPath, Name: handler.Filename})
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
		sendToChunkServer(string(chunk), handler.Filename)
	}

	json.NewEncoder(w).Encode(message{Msg: "File uploaded succesfukly"})
}
func sendToChunkServer(chunk string, fileName string) {
	var chunkName string = generateName()
	request := strings.NewReader(`
		{
			"name":"` + chunkName + `",
			"data":"` + chunk + `"
		}
	`)
	fmt.Println(chunk)
	var chunkNames []string
	chunkNames = append(chunkNames, fileToChunk[fileName]...)
	chunkNames = append(chunkNames, chunkName)
	fileToChunk[fileName] = chunkNames
	serverUrl := chunkServerMap[serverNames[rand.Intn(len(serverNames))]]
	chunkToChunkServer[chunkName] = serverUrl
	resp, err := http.Post(serverUrl+"/upload", "application/json", request)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
