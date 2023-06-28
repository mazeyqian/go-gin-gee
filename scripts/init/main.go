package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
)

// Example: go run scripts/init/main.go -copyData="config.json,database.db"
func main() {
	log.Println("Init...")
	// Define command-line flags
	copyData := flag.String("copyData", "", "Comma-separated list of files to copy from the assets/data directory")
	flag.Parse()
	targetDir := "data"

	// Copy files from the assets/data directory
	if *copyData != "" {
		// Create the target directory if it doesn't exist
		if _, err := os.Stat(targetDir); os.IsNotExist(err) {
			if err := os.Mkdir(targetDir, 0755); err != nil {
				log.Fatalf("Error creating directory %s: %v\n", targetDir, err)
			}
		}

		files := strings.Split(*copyData, ",")
		for _, file := range files {
			src := filepath.Join("assets", targetDir, file)
			dst := filepath.Join(targetDir, file)
			if err := copyFile(src, dst); err != nil {
				log.Printf("Error copying file %s: %v\n", file, err)
			} else {
				log.Printf("Copied file %s\n", file)
			}
		}
	} else {
		script.ListFiles("./assets").ExecForEach("cp -R {{.}} .").Stdout()
	}

	log.Println("All done.")
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
