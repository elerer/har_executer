package main

import (
	"github.com/elerer/hargo"
	"io/ioutil"
	"bufio"
	"bytes"
	"net/url"
)
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/home/elerer/har/dis.har")
	check(err)
	//fmt.Print(string(dat))
	br := bytes.NewReader(dat)
	r := bufio.NewReader(br)
	infurl,_ := url.Parse("http://127.0.0.1:8086")
	hargo.LoadTest("/home/elerer/har/dis.har",r,1,600000000000,*infurl,false)
}
