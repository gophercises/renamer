package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var sep string

func extractKeyVal(path string) (bool, string, string) {
	arr := strings.Split(path, sep)
	if len(arr) == 1 {
		return false, "", ""
	}

	return true, arr[0], arr[1]
}

func mapPaths(path string, names *map[string][]string) error {
	err := filepath.Walk(path, func(pt string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error while waling %q: %v\n", pt, err)
			return err
		}

		if pt == path {
			return nil
		}

		valid, key, value := extractKeyVal(pt)
		if valid {
			use := *names
			if use[key] == nil {
				use[key] = []string{value}
			} else {
				use[key] = append(use[key], value)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return err
	}

	return nil
}

func main() {
	dirPath := flag.String("dir", "./sample/", "path of the directory which needs renaming of files")
	seperator := flag.String("sep", "_", "seperator between filename and number")
	flag.Parse()

	names := make(map[string][]string)
	sep = *seperator

	mapPaths(*dirPath, &names)

	for k, v := range names {
		tot := len(v)

		for i, p := range v {
			oldPath := k + sep + p
			ext := filepath.Ext(p)
			pathSuffix := fmt.Sprintf(" (%d of %d)", i+1, tot)
			newPath := k + pathSuffix + ext
			err := os.Rename(oldPath, newPath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// go run main.go --dir="./sample" --sep="_"
