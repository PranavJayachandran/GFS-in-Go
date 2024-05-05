package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const nameCodeSize = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateName() string {
	rand.Seed(time.Now().UnixNano())
	randomString := make([]byte, nameCodeSize)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}
	return string(randomString)
}

func setHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}

func saveMapsIntoJson() {
	err := os.MkdirAll("backup", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	saveMap("fileToChunk")
	saveMap("chunkToChunkServer")
	saveMap("chunkServerMap")

	jsonData, err := json.MarshalIndent(folderRoot, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	// Save JSON string to file
	err = ioutil.WriteFile("tree.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
func saveMap(fileName string) {
	filePath := filepath.Join("backup", fileName)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creatingg fle ", fileName, err)
		return
	}
	neededMap := getMapFromName(fileName)

	jsonData, err := json.Marshal(neededMap)
	if err != nil {
		fmt.Printf("Error converting %s map to json: %v\n", fileName, err)
		return
	}
	_, err = io.Copy(f, strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error storing fileToChunk map", err)
		return
	}
	fmt.Println(fileName, "map has been stored")
}

func retriveMapsFromJson() {
	retriveMap("fileToChunk")
	retriveMap("chunkToChunkServer")
	retriveMap("chunkServerMap")
}
func retriveMap(fileName string) {
	neededMap := getMapFromName(fileName)
	fileData, err := ioutil.ReadFile("backup/" + fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(fileData, &neededMap)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("Retrived map from file", neededMap)
}
func getMapFromName(mapName string) map[string]interface{} {
	var neededMap map[string]interface{}

	switch mapName {
	case "fileToChunk":
		neededMap = make(map[string]interface{})
		for key, value := range fileToChunk {
			neededMap[key] = value
		}
	case "chunkToChunkServer":
		neededMap = make(map[string]interface{})
		for key, value := range chunkToChunkServer {
			neededMap[key] = value
		}
	case "chunkServerMap":
		neededMap = make(map[string]interface{})
		for key, value := range chunkServerMap {
			neededMap[key] = value
		}
	default:
		fmt.Println("The fileName provided didn't match")
	}
	return neededMap
}
