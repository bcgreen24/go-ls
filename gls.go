package main

import "go-colortext"
import (
	"container/list"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	var option = flag.Bool("size", false, "Show file sizes")
	var mod_date = flag.Bool("date", false, "Show modification dates")
	flag.Parse()
	dir, _ := os.Getwd()
	file_list, _ := filepath.Glob(dir + "/*")
	//var regularFiles = list.New()
	//var dirFiles = list.New()

	var regularFiles = map[string]map[string]string{}
	var dirFiles = map[string]map[string]string{}

	for _, file := range file_list {
		f, _ := os.Open(file)
		fi, _ := f.Stat()
		switch mode := fi.Mode(); {
		case mode.IsDir():
			dirFiles["name"] = map[string]string{
				"size":strconv.FormatInt(fi.Size(), 10),
				"data":
			}
		case mode.IsRegular():
			regularFiles.PushFront(fi.Name())
		}
	}
	for x := dirFiles.Front(); x != nil; x = x.Next() {
		fmt.Println(x.Value)
	}

	for x := regularFiles.Front(); x != nil; x = x.Next() {
		fmt.Println(x.Value)
	}

	for _, file := range file_list {
		f, _ := os.Open(file)
		fi, _ := f.Stat()
		switch mode := fi.Mode(); {
		case mode.IsDir():
			ct.ChangeColor(ct.Green, true, ct.None, false)
			fmt.Printf(filepath.Base(file) + "\n")
		case mode.IsRegular():
			ct.ChangeColor(ct.White, true, ct.None, false)
			if *option == true {
				fmt.Printf(strconv.FormatInt(fi.Size(), 10) + " | ")
			}
			if *mod_date == true {
				const layout = "Jan 2, 2006 at 3:04pm"
				fmt.Printf(fi.ModTime().Format(layout) + " | ")
			}
			fmt.Printf(filepath.Base(file) + "\n")
		}
	}
	ct.ResetColor()
}
