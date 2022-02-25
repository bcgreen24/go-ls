package main

import "go-colortext"
import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

const layout = "Jan 2, 2006 at 3:04pm"

var regularFiles = make(map[string]map[string]string)
var dirFiles = make(map[string]map[string]string)

func main() {

	dir, _ := os.Getwd()
	fileList, err := filepath.Glob(dir + "/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range fileList {
		f, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		switch mode := f.IsDir(); {
		case mode == true:
			dirFiles[f.Name()] = map[string]string{
				"size": strconv.FormatInt(f.Size(), 10),
				"date": f.ModTime().Format(layout),
			}
		case mode == false:
			regularFiles[f.Name()] = map[string]string{
				"size": strconv.FormatInt(f.Size(), 10),
				"date": f.ModTime().Format(layout),
			}
		}
	}

	dirKeys := make([]string, 0, len(dirFiles))
	for k := range dirFiles {
		dirKeys = append(dirKeys, k)
	}

	sort.Strings(dirKeys)

	var fileKeys = make([]string, 0, len(regularFiles))
	for k := range regularFiles {
		fileKeys = append(fileKeys, k)
	}

	sort.Strings(fileKeys)

	for i := range dirKeys {
		ct.ChangeColor(ct.Green, true, ct.None, false)
		fmt.Printf("%v\n", dirKeys[i])
	}

	output := ""

	for i := range fileKeys {
		ct.ChangeColor(ct.White, true, ct.None, false)
		line := regularFiles[fileKeys[i]]
		output += fmt.Sprintf(" %-32s", line["date"])
		output += fmt.Sprintf("%-12s ", line["size"])
		output += fmt.Sprintf(" %v\n", fileKeys[i])

	}
	fmt.Println(output)
	ct.ResetColor()
}
