package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	var root = "sample/"
	var regex, _ = regexp.Compile("(.*)_[0-9]{3}.txt$")
	var filesMap = make(map[string][]string)

	filepath.Walk(root, populateMap(filesMap, regex))

	if len(filesMap) == 0 {
		fmt.Println("Nothing to rename.")
		return
	}

	for _, pathes := range filesMap {
		for i, path := range pathes {
			newPathPattern := fmt.Sprintf("$1 (%d of %d).txt", i+1, len(pathes))
			newPath := regex.ReplaceAllString(path, newPathPattern)
			fmt.Println("Renaming", path, "to", newPath)
			os.Rename(path, newPath)
		}
	}
}

func populateMap(files map[string][]string, regex *regexp.Regexp) func(string, fs.FileInfo, error) error {
	return func(path string, fi fs.FileInfo, err error) error {
		if err != nil { // I have the feeling I should do this, but I'm not sure.
			return err
		}

		if fi.IsDir() {
			return nil
		}

		// The map in the end is going to be something like this:
		// {
		//   "sample/birthday": [
		//     "sample/birthday_001.txt",
		//     "sample/birthday_002.txt",
		//     "sample/birthday_003.txt",
		//     "sample/birthday_004.txt",
		//   ],
		//   "sample/nested/n": [
		//     "sample/nested/n_008.txt",
		//     "sample/nested/n_009.txt",
		//     "sample/nested/n_010.txt",
		//   ]
		// }
		// The keys are not important, they just have to be
		// some kind of identifier so we can count how many
		// of those files are going to get renamed.
		if regex.MatchString(path) {
			fileNameBase := regex.ReplaceAllString(path, "$1")
			files[fileNameBase] = append(files[fileNameBase], path)
		}
		return nil
	}
}
