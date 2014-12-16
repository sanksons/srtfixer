package main

import (
	//"fmt"
	//"os"
	//"strings"
	"srtfixer/srt"
	"flag"
)

type Config struct {
	File       string
	Lag        int64
}

var gConfig Config

func init() {
	filePtr := flag.String("file", "", "a string")
	secPtr  := flag.Int("sec", 0, "a string")
	msecPtr := flag.Int("msec", 0, "a string")
	flag.Parse()
	lag := int64(*msecPtr * 1000000  + *secPtr * 1000000000)
	gConfig = Config {
		File: *filePtr,
		Lag: lag,
	}
}

func main() {
    data, _ := srt.ReadFile(gConfig.File)
	data = srt.AddLagToTime(data, gConfig.Lag)
    srt.WriteFile(data, gConfig.File)
}
