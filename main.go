package main

import (
	"fmt"
	"github.com/xfrr/goffmpeg/transcoder"
)

var inputPath = "./data/input.wav"
var outputPath = "./data/output.pcm"

func main() {

	// Create new instance of transcoder
	trans := new(transcoder.Transcoder)

	// Initialize transcoder passing the input file path and output file path
	err := trans.Initialize(inputPath, outputPath)
	if err != nil {
		fmt.Printf("Initialize error %s", err)
	}
	// Handle error...
	trans.MediaFile().SetAudioChannels(1)
	trans.MediaFile().SetOutputFormat("s16le")

	// Start transcoder process without checking progress
	done := trans.Run(false)

	// This channel is used to wait for the process to end
	err = <-done
	if err != nil {
		fmt.Printf("Trans error %s", err)
	}
	// Handle error...

}
