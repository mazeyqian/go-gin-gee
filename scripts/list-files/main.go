package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os/exec"
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
	files := flag.String("file", "", "Files Formated")        // "file1,file2,file3"
	folders := flag.String("folder", "", "Folders Formated")  // "folder1,folder2,folder3"
	root := flag.String("root", "", "Root of Folders")        // "src"
	esConf := flag.String("esConf", "", "Eslint Config Path") // ".eslintrc.js"
	ext := flag.String("ext", "", "File Ext")                 // ".vue"
	befCom := flag.String("befCom", "", "Before Commands")    // ""
	aftCom := flag.String("aftCom", "", "After Commands")     // ""
	flag.Parse()
	log.Printf("Files: %s", *files)
	log.Printf("Folders: %s", *folders)
	log.Printf("Root: %s", *root)
	log.Printf("ESLint Config Path: %s", *esConf)
	log.Printf("File Ext: %s", *ext)
	log.Printf("Before Commands: %s", *befCom)
	log.Printf("After Commands: %s", *aftCom)
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
	if *befCom != "" {
		befCmd := exec.Command("/bin/sh", "-c", *befCom)
		befResult, err := befCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Before Result: %s", befResult)
	}
	for _, file := range fileArr {
		log.Printf("file: %s", file)
		cmdLines := ""
		if *root != "" {
			cmdLines += fmt.Sprintf("cd %s;", *root)
		}
		cmdLines += fmt.Sprintf("npx eslint %s --fix -c %s;", file, *esConf)
		log.Printf("cmdLines: %s", cmdLines)
		cmd := exec.Command("/bin/sh", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("Result: %s", result)
		count++
	}
	if *aftCom != "" {
		aftCmd := exec.Command("/bin/sh", "-c", *aftCom)
		aftResult, err := aftCmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Printf("After Result: %s", aftResult)
	}
	log.Printf("total: %d", count)
}
