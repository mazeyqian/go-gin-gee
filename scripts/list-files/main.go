package main

import (
	"flag"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func find(root, ext string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func main() {
	count := 0
	files := flag.String("file", "", "Files Formated")       // "file1,file2,file3"
	folders := flag.String("folder", "", "Folders Formated") // "folder1,folder2,folder3"
	esConf := flag.String("esConf", "", "Eslint Config Path")
	ext := flag.String("ext", "", "File Ext") // ".vue"
	flag.Parse()
	log.Printf("Files: %s", *files)
	log.Printf("Folders: %s", *folders)
	log.Printf("ESLint Config Path: %s", *esConf)
	log.Printf("File Ext: %s", *ext)
	fileArr := []string{}
	if *files != "" {
		fileArr = strings.Split(*files, ",")
	}
	folderArr := []string{}
	if *folders != "" {
		folderArr = strings.Split(*folders, ",")
	}
	log.Printf("fileArr Length: %d", len(fileArr))
	log.Printf("folderArr Length: %d", len(folderArr))
	for _, folder := range folderArr {
		log.Printf("folder: %s", folder)
		fileArr = append(fileArr, find(folder, *ext)...)
	}
	for _, file := range fileArr {
		log.Printf("file: %s", file)
		count++
	}
	log.Printf("total: %d", count)
}
