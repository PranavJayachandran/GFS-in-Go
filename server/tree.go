package main

import "fmt"

func setUp() {
	if folderRoot == nil {
		folderRoot = &folderTree{
			folder: []*folderTree{},
			file:   []string{},
		}
	} else {
		folderRoot.folder = []*folderTree{}
		folderRoot.file = []string{}
	}
}

func addFolder(folderData createFileFolderType) {
	var temp = folderRoot
	for _, element := range folderData.Path {
		for _, folderElement := range folderRoot.folder {
			if folderElement.folderName == element {
				folderRoot = folderElement
				break
			}
		}
	}
	var newNode = &folderTree{
		folderName: folderData.Name,
		folder:     []*folderTree{},
		file:       []string{},
	}
	folderRoot.folder = append(folderRoot.folder, newNode)
	folderRoot = temp
	printTree(folderRoot)
}
func addFile(fileData createFileFolderType) {
	var temp = folderRoot
	for _, element := range fileData.Path {
		for _, folderElement := range folderRoot.folder {
			if folderElement.folderName == element {
				folderRoot = folderElement
				break
			}
		}
	}
	folderRoot.file = append(folderRoot.file, fileData.Name)
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
			for _, element := range file.file {
				fmt.Print(element)
			}
			fmt.Print("]\t")
			tempqueue = append(tempqueue, file.folder...)
		}
		fmt.Println()
		queue = tempqueue

	}

}
