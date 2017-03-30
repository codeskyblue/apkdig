package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/shogo82148/androidbinary"
)

func manifest() {
	f, _ := os.Open("AndroidManifest.xml")
	xml, err := androidbinary.NewXMLFile(f)
	if err != nil {
		log.Fatal(err)
	}
	reader := xml.Reader()
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(data))
	println("Hello")
}

func tableFile() {
	f, _ := os.Open("resources.arsc")
	tf, err := androidbinary.NewTableFile(f)
	if err != nil {
		log.Fatal(err)
	}
	config := &androidbinary.ResTableConfig{}
	resId := androidbinary.ResId(0x7F060002)
	val, err := tf.GetResource(resId, config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v %v", val, config)
}

func main() {
	tableFile()
}
