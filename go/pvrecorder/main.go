package main

import (
	. "github.com/Picovoice/pvrecorder/binding/go"
	"log"
)

func main() {
	recorder := NewPvRecorder( /*FrameLength*/ 512)
	err := recorder.Init()
	if err != nil {
		// handle init error
	}
	defer recorder.Delete()

	err = recorder.Start()
	if err != nil {
		// handle start error
	}

	frame, err := recorder.Read()
	if err != nil {
		log.Fatal(err)
	}

	_ = frame

	err = recorder.Stop()
	if err != nil {
		log.Fatal(err)
	}

	recorder.Delete()

}
