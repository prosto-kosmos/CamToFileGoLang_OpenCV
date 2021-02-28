package main

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	var filesDir string = "files"
	for {
		if _, err := os.Stat(filesDir); !os.IsNotExist(err) {
			files, err := ioutil.ReadDir(filesDir)
			if err != nil {
				log.Fatal(err)
			}

			if len(files) != 0 {
				println("sending: ", files[len(files)-1].Name())
				os.Remove(filesDir + "/" + files[len(files)-1].Name())
			} else {
				time.Sleep(5 * time.Second)
				println("no files")
			}
		} else {
			time.Sleep(5 * time.Second)
			println("no folder")
		}
	}
}
