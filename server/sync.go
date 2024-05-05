package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

var serverNames []string

func heartBeat() {
	for key, value := range chunkServerMap {
		fmt.Println("Heartbeat call to ", key, value)
		res, err := http.Get(value + "/heartBeat")
		if err != nil {
			fmt.Println(err.Error())
		}
		scanner := bufio.NewScanner(res.Body)
		for i := 0; scanner.Scan(); i++ {
			fmt.Println(scanner.Text())
		}
	}

}
func attendance(w http.ResponseWriter, req *http.Request) {
	x := &chunkServer{}
	fmt.Println("Called")
	err := json.NewDecoder(req.Body).Decode(x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("got user:", x)
	serverNames = append(serverNames, x.Id)
	chunkServerMap[x.Id] = x.Url
	fmt.Fprintf(w, "hello\n")
}
