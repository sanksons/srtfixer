package main

//import the srt package which contains necessary SRT reading writing functions
//import the flag package to handle command line arguments
import (
	"srtfixer/srt"
	"flag"
)

type Config struct {
	File       string
	Lag        int64
}

var gConfig Config

func init() {
	filePtr := flag.String("file", "", "Path to srt file")
	secPtr  := flag.Int("sec", 0, "Seconds to RUSH or LAG")
	msecPtr := flag.Int("msec", 0, "microseconds to RUSH or LAG")
	//parse the cli flags
	flag.Parse()
	//calculate lag by combining seconds and microseconds
	lag := int64(*msecPtr * 1000000  + *secPtr * 1000000000)
	//setup params in global config
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
