# Program for recording files from webcam and sending them to Google Drive in Go language

This application contains a module for analyzing and processing a video stream from a webcam (using OpenCV) and saving it to files of a certain duration, as well as modules for working with the Google Drive storage, including sending these files to it. The advantage is the ability to cross compile and run on single board computers such as Orange Pi and Raspberry Pi

Это приложение содержит модуль анализа и обработки видеопотока с вебкамеры (используя OpenCV) и сохраниния его в файлы определенной длительности, а также модули работы с хранилищем Google Drive, в том числе отправки этих файлов в него. Приимуществом является возможность кросскомпиляции и запуска на одноплатных компьютерах, таких как Orange Pi и Raspberry Pi

## How to use 

First, you need to install Go on your computer. Installation example on Ubuntu [here](https://losst.ru/ustanovka-go-ubuntu)

Then you need to install [GoCV](https://gocv.io/getting-started/linux/)

Everything is now ready to run. Follow to start webcam recording:

    go run record.go

To send files read [this](https://developers.google.com/drive/api/v3/quickstart/go) and do:

    go run save saveFile.go

## Links
[An article on using GoCV](https://habr.com/ru/company/skillbox/blog/462159/)
[Article with examples of cross-compilation](https://habr.com/ru/post/249449/)