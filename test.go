package main

import (
	"time"
)

// // WriteVideo in file
// func WriteVideo() {
// 	time.Sleep(10 * time.Second)
// 	println("file was written")
// }

// // SaveFile to Disk
// func SaveFile() {
// 	time.Sleep(15 * time.Second)
// 	println("file was saved")
// }

// func main() {
// 	theMine := [5]string{"ore1", "ore2", "ore3"}
// 	oreChan := make(chan string)

// 	// WriteVideo in file
// 	go func(mine [5]string) {
// 		time.Sleep(10 * time.Second)
// 		for _, item := range mine {
// 			oreChan <- item //отправка
// 		}
// 		println("file was written")
// 	}(theMine)

// 	// SaveFile to Disk
// 	go func() {
// 		// time.Sleep(15 * time.Second)
// 		for i := 0; i < 3; i++ {
// 			foundOre := <-oreChan // получение
// 			fmt.Println("Miner: Received " + foundOre + " from finder")
// 		}
// 		println("file was saved")
// 	}()

// 	/*var filesDir string = "files"
// 	for {
// 		if _, err := os.Stat(filesDir); !os.IsNotExist(err) {
// 			files, err := ioutil.ReadDir(filesDir)
// 			if err != nil {
// 				log.Fatal(err)
// 			}

// 			if len(files) != 0 {
// 				println("sending: ", files[len(files)-1].Name())
// 				os.Remove(filesDir + "/" + files[len(files)-1].Name())
// 			} else {
// 				time.Sleep(5 * time.Second)
// 				println("no files")
// 			}
// 		} else {
// 			time.Sleep(5 * time.Second)
// 			println("no folder")
// 		}
// 	}*/
// }

func main() {
	for {
		time.Sleep(5 * time.Second)
		println("file was written", time.Now().String())
		go Sending("file name")
	}

}

// Sending file
func Sending(fn string) {
	time.Sleep(10 * time.Second)
	println("file was saved", time.Now().String())
}
