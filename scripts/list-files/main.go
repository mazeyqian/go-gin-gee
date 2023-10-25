package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/samber/lo"
)

func find(absRoot, root, ext string) []string {
	var a []string
	if absRoot != "" {
		root = absRoot + root
	}
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		// log.Printf("find s: %s", s)
		if filepath.Ext(d.Name()) == ext {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func main() {
	count := 0
	files := flag.String("files", "", "Files Formated")       // "file1,file2,file3"
	folders := flag.String("folders", "", "Folders Formated") // "folder1,folder2,folder3"
	root := flag.String("root", "", "Root of Folders")        // "src"
	esConf := flag.String("esConf", "", "Eslint Config Path") // ".eslintrc.js"
	ext := flag.String("ext", "", "File Ext")                 // ".vue"
	filesRang := flag.String("filesRang", "", "Files Range")  // ""
	befCom := flag.String("befCom", "", "Before Commands")    // ""
	aftCom := flag.String("aftCom", "", "After Commands")     // ""
	flag.Parse()
	log.Printf("Files: %s", *files)
	log.Printf("Folders: %s", *folders)
	log.Printf("Root: %s", *root)
	log.Printf("ESLint Config Path: %s", *esConf)
	log.Printf("File Ext: %s", *ext)
	log.Printf("Files Range: %s", *filesRang)
	log.Printf("Before Commands: %s", *befCom)
	log.Printf("After Commands: %s", *aftCom)
	fileArr := []string{}
	if *files != "" {
		fileArr = strings.Split(*files, ",")
		if *root != "" {
			for i, file := range fileArr {
				fileArr[i] = *root + file
			}
		}
	}
	folderArr := []string{}
	if *folders != "" {
		folderArr = strings.Split(*folders, ",")
	}
	// log.Printf("fileArr Length: %d", len(fileArr))
	// log.Printf("folderArr Length: %d", len(folderArr))
	for _, folder := range folderArr {
		// log.Printf("folder: %s", folder)
		fileArr = append(fileArr, find(*root, folder, *ext)...)
	}
	// log.Printf("fileArr Length: %d", len(fileArr))
	rootCom := ""
	if *root != "" {
		rootCom = fmt.Sprintf("cd %s;", *root)
	}
	if *befCom != "" {
		befStr := fmt.Sprintf("%s%s", rootCom, *befCom)
		// log.Printf("befStr: %s", befStr)
		befCmd := exec.Command("/bin/bash", "-c", befStr) // *befCom)
		_, err := befCmd.CombinedOutput()
		if err != nil {
			log.Println("Shell Error:", err)
		}
		// fmt.Printf("Before Result: %s", befResult)
	}
	for _, file := range fileArr {
		log.Printf("File: %s", file)
		cmdLines := ""
		if *root != "" {
			cmdLines += fmt.Sprintf("cd %s;", *root)
		}
		cmdLines += fmt.Sprintf("npx eslint %s --fix -c %s;", file, *esConf)
		// cmdLines += fmt.Sprintf("exit %d;", 0)
		// log.Printf("cmdLines: %s", cmdLines)
		cmd := exec.Command("/bin/bash", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("Eslint Error:", err)
		}
		resultStr := string(result)
		if resultStr == "" {
			resultStr = "ok"
		}
		log.Printf("Eslint Result: %s", resultStr)
		count++
	}
	if *aftCom != "" {
		aftStr := fmt.Sprintf("%s%s", rootCom, *aftCom)
		// log.Printf("aftStr: %s", aftStr)
		aftCmd := exec.Command("/bin/bash", "-c", aftStr)
		_, err := aftCmd.CombinedOutput()
		if err != nil {
			log.Println("Shell Error:", err)
		}
		// log.Printf("After Result: %s", aftResult)
	}
	log.Printf("Worked Count: %d", count)
	otherFiles := []string{}
	totalFiles := []string{}
	if *filesRang != "" {
		totalFiles = find("", *filesRang, *ext)
		// for _, file := range totalFiles {
		// 	log.Printf("Total File: %s", file)
		// }
		log.Printf("Total Count: %d", len(totalFiles))
		otherFiles = lo.Filter(totalFiles, func(s string, index int) bool {
			return !lo.Contains(fileArr, s)
		})
		log.Printf("Other Count: %d", len(otherFiles))
	}
	for _, file := range fileArr {
		log.Printf("Worked File: %s", file)
	}
	for _, file := range otherFiles {
		log.Printf("Other File: %s", file)
	}
	for _, file := range totalFiles {
		log.Printf("Total File: %s", file)
	}
}
