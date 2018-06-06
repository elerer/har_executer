package main

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/elerer/hargo"
	"io/ioutil"
	"net/url"
	"strings"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	modePtr := flag.String("mode", "har", "har - follows har file and validate | load - perform load test")
	filePtr := flag.String("hf", "/home/elerer/har/dis.har", "path to har file")
	flag.Parse()

	isHar := strings.Compare(*modePtr, "har") == 0
	dat, err := ioutil.ReadFile(*filePtr)
	check(err)
	fmt.Printf("run %s test , har file %s\n",*modePtr,*filePtr)
	//fmt.Print(string(dat))
	br := bytes.NewReader(dat)
	r := bufio.NewReader(br)
	infurl, _ := url.Parse("http://127.0.0.1:8086")
	hargo.LoadTest("/home/elerer/har/dis.har", r, 1, 600000000000, *infurl, false,isHar)
}
