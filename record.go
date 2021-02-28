package main

import (
	"image"
	"log"
	"os"
	"time"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("error opening device: %v", err)
	}
	defer webcam.Close()

	img := gocv.NewMat()
	defer img.Close()

	window := gocv.NewWindow("WebCamWindow")
	defer window.Close()

	var frameWidth int = int(webcam.Get(gocv.VideoCaptureFrameWidth))
	var frameHeight int = int(webcam.Get(gocv.VideoCaptureFrameHeight))
	var fps float64 = webcam.Get(gocv.VideoCaptureFPS)
	var codec string = "MJPG"

	var durationOneFileSeconds int = 5 // количество секунд видео в одном файле
	var filesDir string = "files"
	var prefixPath string = filesDir + "/video"
	var mediaContainer string = ".mkv"
	var frameInOneFileCounter int = 0 // тупо счетчик

	if _, err := os.Stat(filesDir); os.IsNotExist(err) {
		err = os.Mkdir(filesDir, 0777)
		if err != nil {
			log.Fatalf("error create folder: '%s'", err)
		}
	}

	fileName := prefixPath + time.Now().Format("_2006-01-02_15:04:05") + mediaContainer
	videoFile, err := gocv.VideoWriterFile(fileName, codec, fps, frameWidth, frameHeight, true)
	if err != nil {
		log.Fatalf("error createing VideoWriterFile: %v", err)
	}

	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Println("Unable to read from the webcam")
			continue
		}

		// обработка

		if int(fps)*durationOneFileSeconds < frameInOneFileCounter {
			fileName = prefixPath + time.Now().Format("_2006-01-02_15:04:05") + mediaContainer
			videoFile, err = gocv.VideoWriterFile(fileName, codec, fps, frameWidth, frameHeight, true)
			frameInOneFileCounter = 0
		}
		videoFile.Write(img)
		frameInOneFileCounter++

		gocv.Resize(img, &img, image.Point{960, 540}, 100, 100, 1)
		window.IMShow(img)
		c := window.WaitKey(50)
		if c == 27 {
			break
		}
	}
}
