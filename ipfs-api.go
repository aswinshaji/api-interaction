package main

import (
	"fmt"
	"io/ioutil"
    "log"
	"bytes"
	"github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell


func main() {
	sh = shell.NewShell("localhost:5001")

	// Fetching data from IPFS
	err := sh.Get("QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9", "data.txt")
	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}

	text := string(content)
    fmt.Println(text)
	
	fmt.Println("Successfully read the sample data..")

	// Storing string data into IPFS
	cid, err := sh.Add(bytes.NewBufferString("Hello world"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cid)
	fmt.Println("Successfully added Hello World string to the IPFS..")

	// Fetching the data with CID
	err = sh.Get(cid, "hello.txt")
	content2, err := ioutil.ReadFile("hello.txt")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}

	text2 := string(content2)
    fmt.Println(text2)

	fmt.Println("Successfully fetched back the Hello World string and stored to hello.txt..")
	fmt.Println("All Good!!")
}