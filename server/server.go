package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var folderRoot *folderTree

func getFolderStructure(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	data := &pathReq{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	folder := folderRoot
	fmt.Println(data, folderRoot.folders)
	for _, element := range data.Path {
		for _, folderElement := range folder.folders {
			if folderElement.folderName == element {
				folder = folderElement
				break
			}
		}
	}
	var folderNames = []string{}
	for _, element := range folder.folders {
		folderNames = append(folderNames, element.folderName)
	}
	json.NewEncoder(w).Encode(folderStructure{Folders: folderNames, Files: folder.files})
}
func createFolder(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	folderData := &createFileFolderType{}
	err := json.NewDecoder(r.Body).Decode(folderData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("got user:", folderData.Path)
	addFolder(*folderData)
}

func main() {
	go setUp()
	// retriveMapsFromJson()
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "X-Requested-With"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
	r := mux.NewRouter()
	r.HandleFunc("/attendance", attendance)
	r.HandleFunc("/folderStructure", getFolderStructure)
	r.HandleFunc("/upload", uploadFile)
	r.HandleFunc("/createFolder", createFolder)
	http.ListenAndServe(":8080", cors(r))

}
