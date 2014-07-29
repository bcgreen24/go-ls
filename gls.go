package main

import "go-colortext"
import (
	"flag"
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
	var option = flag.Bool("size", false, "Show file sizes")
	var mod_date = flag.Bool("date", false, "Show modification dates")

	flag.Parse()
	dir, _ := os.Getwd()
	file_list, err := filepath.Glob(dir + "/*")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range file_list {
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

	dir_keys := make([]string, 0, len(dirFiles))
	for k := range dirFiles {
		dir_keys = append(dir_keys, k)
	}

	sort.Strings(dir_keys)

	file_keys := make([]string, 0, len(regularFiles))
	for k := range regularFiles {
		file_keys = append(file_keys, k)
	}

	sort.Strings(file_keys)

	for i := range dir_keys {
		ct.ChangeColor(ct.Green, true, ct.None, false)
		fmt.Println(dir_keys[i])
	}

	for i := range file_keys {
		ct.ChangeColor(ct.White, true, ct.None, false)
		line := regularFiles[file_keys[i]]
		output := file_keys[i]
		if *option == true {
			output += " | " + line["size"]
		}
		if *mod_date == true {
			output += " | " + line["date"]
		}
		fmt.Println(output)
	}
	ct.ResetColor()
}
