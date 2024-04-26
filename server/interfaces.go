package main

const (
	byteSize = 10
)

type chunkServer struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}
type pathReq struct {
	Path string `json:"path"`
}

type folderStructure struct {
	Folders []string `json:"folders"`
	Files   []string `json:"files"`
}

type message struct {
	Msg string `json:"msg"`
}
