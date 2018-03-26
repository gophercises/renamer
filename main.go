package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// fileName := "birthday_001.txt"
	// // => Birthday - 1 of 4.txt
	// newName, err := match(fileName, 4)
	// if err != nil {
	// 	fmt.Println("no match")
	// 	os.Exit(1)
	// }
	// fmt.Println(newName)
	dir := "./sample"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	count := 0
	var toRename []string
	for _, file := range files {
		if file.IsDir() {
		} else {
			_, err := match(file.Name(), 0)
			if err == nil {
				count++
				toRename = append(toRename, file.Name())
			}
		}
	}
	for _, origFilename := range toRename {
		origPath := filepath.Join(dir, origFilename)
		newFilename, err := match(origFilename, count)
		if err != nil {
			panic(err)
		}
		newPath := filepath.Join(dir, newFilename)
		fmt.Printf("mv %s => %s\n", origPath, newPath)
		err = os.Rename(origPath, newPath)
		if err != nil {
			panic(err)
		}
	}
}

// match returns the new file name, or an error if the file name
// didn't match our pattern.
func match(filename string, total int) (string, error) {
	// "birthday", "001", "txt"
	pieces := strings.Split(filename, ".")
	ext := pieces[len(pieces)-1]
	tmp := strings.Join(pieces[0:len(pieces)-1], ".")
	pieces = strings.Split(tmp, "_")
	name := strings.Join(pieces[0:len(pieces)-1], "_")
	number, err := strconv.Atoi(pieces[len(pieces)-1])
	if err != nil {
		return "", fmt.Errorf("%s didn't match our pattern", filename)
	}
	// Birthday - 1.txt
	return fmt.Sprintf("%s - %d of %d.%s", strings.Title(name), number, total, ext), nil
}
