package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/elerer/hargo"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getHarsPath(arg *string) string {
	var harPath string
	//Get process executable path, for relative hars dir
	if strings.Compare(*arg, "proc") == 0 {
		pwd, _ := os.Executable()
		exPath := filepath.Dir(pwd)
		harPath = exPath + "/hars"
	} else {
		harPath = *arg
	}

	return (harPath + "/")
}

func getHarsFileInfo(fn, harPath *string) []os.FileInfo {
	var files []os.FileInfo

	if strings.Compare(*fn, "all") == 0 {
		files, _ = ioutil.ReadDir(*harPath)
	} else { //get fileinfo for single file
		filePath := *harPath + *fn
		fs, _ := os.Create(filePath)
		stat, _ := fs.Stat()
		files = append(files, stat)
	}

	return files
}

func main() {
	modePtr := flag.String("mode", "har", "har - follows har file and validate | load - perform load test")
	fileName := flag.String("hf", "all", "har file name, 'all' for running all hars in ./hars ")
	harsPath := flag.String("path", "proc", "path to hars")
	dur := flag.Int("dur", 10000, "duration in ms")
	workers := flag.Int("workers", 1, "How many parrallel workers")

	flag.Parse()

	harPath := getHarsPath(harsPath)

	fmt.Printf("getting hars from %s\n", harPath)
	//which mode to run test har/load
	isHar := strings.Compare(*modePtr, "har") == 0

	files := getHarsFileInfo(fileName, &harPath)

	//insert har files to run into slice and then iterate and run

	//Time out should be enforced here,
	d := time.Duration(*dur)
	dd := d * time.Millisecond
	timeout := time.After(dd)

	//Limit num of workers
	fmt.Printf("will spawn %d work in parrallel", *workers)
	ch := make(chan int, *workers)

	for {
		for _, file := range files {
			fmt.Printf("file %s\n", file.Name())
			if strings.HasSuffix(file.Name(), "har") {
				fmt.Printf("reading %s\n", file.Name())
				dat, _ := ioutil.ReadFile(harPath + file.Name())

				br := bytes.NewReader(dat)
				r := bufio.NewReader(br)

				har, _ := hargo.Decode(r)

				nonUrl, _ := url.Parse("http://nonnononon.com")

				ch <- 0

				go hargo.LoadTest(*fileName, har, *nonUrl, false, isHar, ch)

			}

		}

		if isHar == true {
			return
		}

		select {
		case <-timeout:
			fmt.Println("test duration passed...")
			return
		default:

		}

	}

}
