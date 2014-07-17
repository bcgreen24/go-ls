package main

import "go-colortext"
import (
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
	for _, file := range file_list {
		f, _ := os.Open(file)
		fi, _ := f.Stat()
		
		switch mode := fi.Mode(); {
		case mode.IsDir():
			ct.ChangeColor(ct.Green, true, ct.None, false)
			fmt.Printf(file + "\n")
		case mode.IsRegular():
			ct.ChangeColor(ct.White, true, ct.None, false)
			if *option == true {
				fmt.Printf(strconv.FormatInt(fi.Size(), 10) + " ")
			}
			if *mod_date == true {
				const layout = "Jan 2, 2006 at 3:04pm"
				fmt.Printf(fi.ModTime().Format(layout) + " ")
			}
			fmt.Printf(file + "\n")
		}
	}
	ct.ResetColor()
}
