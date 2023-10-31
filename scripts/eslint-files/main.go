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
	exts := []string{}
	if ext != "" {
		exts = strings.Split(ext, ",")
	}
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if ext == "" {
			a = append(a, s)
		} else {
			for _, ext := range exts {
				if filepath.Ext(d.Name()) == ext {
					a = append(a, s)
					break
				}
			}
		}
		//  else if filepath.Ext(d.Name()) == ext {
		// 	a = append(a, s)
		// }
		return nil
	})
	return a
}

func main() {
	count := 0
	files := flag.String("files", "", "Files Formated")       // "file1,file2,file3"
	folders := flag.String("folders", "", "Folders Formated") // "folder1,folder2,folder3"
	esConf := flag.String("esConf", "", "ESLint Config Path") // ".eslintrc.js"
	esCom := flag.String("esCom", "", "ESLint Command")       // "--fix"
	root := flag.String("root", "", "Root of Folders")        // "src"
	ext := flag.String("ext", ".js", "File Ext")              // ".vue,.js,.ts,.jsx,.tsx"
	befCom := flag.String("befCom", "", "Before Commands")    // ""
	aftCom := flag.String("aftCom", "", "After Commands")     // ""
	filesRang := flag.String("filesRang", "", "Files Range")  // ""
	flag.Parse()
	log.Printf("Files: %s", *files)
	log.Printf("Folders: %s", *folders)
	log.Printf("Root: %s", *root)
	log.Printf("ESLint Config Path: %s", *esConf)
	log.Printf("ESLint Command: %s", *esCom)
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
	for _, folder := range folderArr {
		fileArr = append(fileArr, find(*root, folder, *ext)...)
	}
	rootCom := ""
	if *root != "" {
		rootCom = fmt.Sprintf("cd %s;", *root)
	}
	if *befCom != "" {
		befStr := fmt.Sprintf("%s%s", rootCom, *befCom)
		befCmd := exec.Command("/bin/bash", "-c", befStr)
		result, err := befCmd.CombinedOutput()
		if err != nil {
			log.Println("Bash Error:", err)
		}
		resultStr := string(result)
		if resultStr == "" {
			resultStr = "ok"
		}
		log.Printf("Bash Result: %s", resultStr)
	}
	for _, file := range fileArr {
		log.Printf("File: %s", file)
		cmdLines := ""
		if *root != "" {
			cmdLines += fmt.Sprintf("cd %s;", *root)
		}
		esConfCom := ""
		if *esConf != "" {
			esConfCom = fmt.Sprintf(" -c %s", *esConf)
		}
		esComCom := ""
		if *esCom != "" {
			esComCom = fmt.Sprintf(" %s", *esCom)
		}
		cmdLines += fmt.Sprintf("npx eslint %s%s%s;", file, esComCom, esConfCom)
		// log.Printf("ESLint Command: %s", cmdLines)
		cmd := exec.Command("/bin/bash", "-c", cmdLines)
		result, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("ESLint Error:", err)
		}
		resultStr := string(result)
		if resultStr == "" {
			resultStr = "ok"
		}
		log.Printf("ESLint Result: %s", resultStr)
		count++
	}
	if *aftCom != "" {
		aftStr := fmt.Sprintf("%s%s", rootCom, *aftCom)
		aftCmd := exec.Command("/bin/bash", "-c", aftStr)
		result, err := aftCmd.CombinedOutput()
		if err != nil {
			log.Println("Bash Error:", err)
		}
		resultStr := string(result)
		if resultStr == "" {
			resultStr = "ok"
		}
		log.Printf("Bash Result: %s", resultStr)
	}
	if *filesRang != "" {
		totalFiles := find("", *filesRang, *ext)
		otherFiles := lo.Filter(totalFiles, func(s string, index int) bool {
			return !lo.Contains(fileArr, s)
		})
		log.Printf("Worked Count: %d", count)
		log.Printf("Other Count: %d", len(otherFiles))
		log.Printf("Total Count: %d", len(totalFiles))
		for _, file := range fileArr {
			log.Printf("Worked File: %s", file)
		}
		for _, file := range otherFiles {
			log.Printf("Other File: %s", file)
		}
		// for _, file := range totalFiles {
		// 	log.Printf("Total File: %s", file)
		// }
	} else {
		log.Printf("Worked Count: %d", count)
	}
}
