package main

import "go-colortext"
import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const layout = "Jan 2, 2006 at 3:04pm"

func main() {
	var option = flag.Bool("size", false, "Show file sizes")
	var mod_date = flag.Bool("date", false, "Show modification dates")

	flag.Parse()
	dir, _ := os.Getwd()
	file_list, err := filepath.Glob(dir + "/*")
	if err != nil {
		log.Fatal(err)
	}
	var regularFiles = make(map[string]map[string]string)
	var dirFiles = make(map[string]map[string]string)

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

	for key, _ := range dirFiles {
		ct.ChangeColor(ct.Green, true, ct.None, false)
		fmt.Println(key)
	}

	for key, value := range regularFiles {
		ct.ChangeColor(ct.White, true, ct.None, false)
		line := key
		if *option == true {
			line += " | " + value["size"]
		}
		if *mod_date == true {
			line += " | " + value["date"]
		}
		fmt.Println(line)
	}
	ct.ResetColor()
}
