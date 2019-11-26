package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"math"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "Print s")
)

func hrSize(fsize int64) string {
	lsize := float64(fsize)
	for lsize > 1 {
		lsize /= 1024
	}
	math.Ceil(lsize)
	fsize = int64(lsize)
	siz := string(fsize) + "KB"
	return siz
}

func printAll(file os.FileInfo, flag bool) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize := strconv.Itoa(int(file.Size()))
	
	if flag {
		lSize := int(file.Size())
		if lSize > 1023 {
			lSize /= 1024
			fSize = strconv.Itoa(lSize) + "KB"
		} else {
			fSize = strconv.Itoa(lSize) + "B"
		}
	}
	fmt.Printf("%s %s %s \n", fSize, time, file.Name())
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)

	for _, file := range files {
		var flag bool = false
		if *h {
			flag = true
		}
		if *a{
			printAll(file, flag)
		} else {
			fmt.Println(file.Name())
		}
	}
}
