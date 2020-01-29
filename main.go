package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "Print s")
	date = flag.Bool("sort_date", false, "Sorted by date")
	size = flag.Bool("sort_size", false, "Sorted by size")
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

type SortByDate []os.FileInfo


func (ss SortByDate) Len() int {
	return len(ss)
}

func (ss SortByDate) Swap(i int, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss SortByDate) Less(i int, j int) bool {
	return ss[i].ModTime().UnixNano() < ss[j].ModTime().UnixNano()
}

type SortBySize []os.FileInfo


func (ss SortBySize) Len() int {
	return len(ss)
}

func (ss SortBySize) Swap(i int, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss SortBySize) Less(i int, j int) bool {
	return ss[i].Size() < ss[j].Size()
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
		if *date {
			var sortBy sort.Interface
			sortBy = SortByDate(files)
			sort.Sort(sortBy)
		}
		if *size {
			var sortBy sort.Interface
			sortBy = SortBySize(files)
			sort.Sort(sortBy)
		}
	}
}
