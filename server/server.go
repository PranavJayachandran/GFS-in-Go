package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getFolderStructure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	data := &pathReq{}
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("got user:", data.Path)
	folderStructure := folderStructure{Folders: []string{data.Path + "one"}, Files: []string{data.Path + "two", data.Path + "three"}}
	json.NewEncoder(w).Encode(folderStructure)
}

func main() {

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Origin", "X-Requested-With"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
	r := mux.NewRouter()
	r.HandleFunc("/attendance", attendance)
	r.HandleFunc("/folderStructure", getFolderStructure)
	r.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", cors(r))

}
