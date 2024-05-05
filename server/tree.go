package main

import (
	"fmt"
	"time"
)

func setUp() {
	if folderRoot == nil {
		folderRoot = &folderTree{
			folders: []*folderTree{},
			files:   []string{},
		}
	} else {
		folderRoot.folders = []*folderTree{}
		folderRoot.files = []string{}
	}

	saveBackupTicker := time.NewTicker(1 * time.Minute)
	defer saveBackupTicker.Stop()

	heartBeatTikcer := time.NewTicker(2 * time.Minute)

	for {
		select {
		case <-saveBackupTicker.C:
			saveMapsIntoJson()
		case <-heartBeatTikcer.C:
			heartBeat()
		}
	}
}

func addFolder(folderData createFileFolderType) {
	var temp = folderRoot
	for _, element := range folderData.Path {
		for _, folderElement := range folderRoot.folders {
			if folderElement.folderName == element {
				folderRoot = folderElement
				break
			}
		}
	}
	var newNode = &folderTree{
		folderName: folderData.Name,
		folders:    []*folderTree{},
		files:      []string{},
	}
	folderRoot.folders = append(folderRoot.folders, newNode)
	folderRoot = temp
	printTree(folderRoot)
}
func addFile(fileData createFileFolderType) {
	var temp = folderRoot
	for _, element := range fileData.Path {
		for _, folderElement := range folderRoot.folders {
			if folderElement.folderName == element {
				folderRoot = folderElement
				break
			}
		}
	}
	folderRoot.files = append(folderRoot.files, fileData.Name)
	folderRoot = temp
	printTree(folderRoot)
}
func printTree(root *folderTree) {
	if root == nil {
		return
	}
	queue := []*folderTree{root}
	for len(queue) > 0 {
		tempqueue := []*folderTree{}
		for _, file := range queue {
			fmt.Print("Folder Name:", file.folderName, "\t[")
			for _, element := range file.files {
				fmt.Print(element)
			}
			fmt.Print("]\t")
			tempqueue = append(tempqueue, file.folders...)
		}
		fmt.Println()
		queue = tempqueue

	}

}
