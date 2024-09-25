package core

import (
	"fmt"

	"gocv.io/x/gocv"
)

func Read(url interface{}) {
	// parse args
	deviceID := "rtsp://admin:123456789a@192.168.0.110:554/Streaming/Channels/101"
	//webcam, err := gocv.VideoCaptureFile("1.mp4")
	webcam, err := gocv.OpenVideoCapture(url)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", url)
		return
	}
	width := webcam.Get(gocv.VideoCaptureFrameWidth)
	height := webcam.Get(gocv.VideoCaptureFrameHeight)
	fmt.Println(width, height)
	defer webcam.Close()
	img := gocv.NewMat()
	defer img.Close()
	for {
		if ok := webcam.Read(&img); !ok {
			webcam.Close()
			webcam, _ = gocv.OpenVideoCapture(url)
			fmt.Printf("Device closed: %v\n", deviceID)
			continue
		}
		if img.Empty() {
			continue
		}
	}
}
