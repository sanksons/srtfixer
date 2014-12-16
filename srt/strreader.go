package srt

//import the required Packages.
import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// Each portion of the SRT such as:
//
// --------- Block 1 -----------------
// 324
// 01:58:11,142 --> 01:58:12,940
// What are you doing?
// ---------- Block 2 ---------------
// 325
// 01:58:49,139 --> 01:58:50,812
// Shut up!
//
// are assumed to be blocks and SRT consists of several of such blocks.
// Further, each block consists of following entities:
// a:) Serial number [starting from 1, mainly]
// b:) Duration for which the subtitle needs to be shown
//     [Start and End time are seperated via "-->" characters]
// c:) Contents to be displayed during this period.
type Block struct {
	Serial  int
	Start   time.Time
	End     time.Time
	Content string
}

//Add time lag.
func (b *Block) AddTimeLag(lag time.Duration) *Block {
	b.Start = b.Start.Add(lag)
	b.End = b.End.Add(lag)
	return b
}

//prepare data based on lag
func AddLagToTime(originalData []Block, lag int64) []Block {

	lagTime := time.Duration(lag)
	for i, _ := range originalData {
		originalData[i].AddTimeLag(lagTime)
	}
	return originalData
}

//Read the file specified by the path.
//Break the SRT file contents into chunks of []Block
func ReadFile(filePath string) ([]Block, error) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []Block{}, err
	}
	blocks := constructBlocks(string(fileContents))
	return blocks, nil
}

//Write the supplied []Block array to specified file.
func WriteFile(blocks []Block, filepath string) bool {
	//open the file in read-write mode.
	f, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		//error found in opening file.
		return false
	}
	//Note: defer is called after function returns.
	defer f.Close()

	dataStr := blocks2String(blocks)
	_, err   = f.Write([]byte(dataStr))
	if err != nil {
		return false
	}
	return true
}

//Extend the []Block array
func extendArray(arr []Block) []Block {
	arrNew := make([]Block, len(arr)+1)
	for i := range arr {
		arrNew[i] = arr[i]
	}
	arr = arrNew
	return arr
}

//Extract time out of the supplied string
//Assumed string to be of format:
// 01:58:11,142 --> 01:58:12,940
func extractTime(rawTime string, tType string) time.Time {
	index := 0
	if tType != "START" {
		index = 1
	}
	aTime := strings.Split(rawTime, " --> ")
	aTime = strings.Split(aTime[index], ",")

	parsedTime, _ := time.Parse("15:04:05", aTime[0])
	//Now we have the time.Time in parsedTime, but we need to add
	//microseconds to it as well.
	microSeconds, _ := strconv.Atoi(aTime[1])
	nanoSeconds := time.Duration(microSeconds) * 1000000
	parsedTime = parsedTime.Add(nanoSeconds)
	return parsedTime
}

//Converts the supplied []Block array to string format.
func blocks2String(blocks []Block) string {
	dataStr := ""
	for _, e := range blocks {
		dataStr += strconv.Itoa(e.Serial) + "\r\n"
		milliSecString := strconv.Itoa(e.Start.Nanosecond() / 1000000)
		milliSecStringE := strconv.Itoa(e.End.Nanosecond() / 1000000)
		dataStr = dataStr + e.Start.Format("15:04:05") + "," + milliSecString +
			" --> " + e.End.Format("15:04:05") + "," + milliSecStringE + "\r\n"
		dataStr = dataStr + e.Content + "\r\n\r\n"
	}
	return dataStr
}

//Construct the []Block array based on the string data supplied.
func constructBlocks(fileContent string) []Block {
	//Remove the carriage return \r characters from content.
	fileContent = strings.Replace(fileContent, "\r", "", -1)
	//Split into blocks containing three sets
	rawCBlocks := strings.Split(fileContent, "\n\n")
	//declare a Slice to store Blocks.
	var blocks []Block
	for _, e := range rawCBlocks {
		block := strings.SplitN(e, "\n", 3)
		if len(block) >= 3 {
			//Extend our container array
			blocks = extendArray(blocks)
			//Extract Serial, Start-End date and Content
			SNo, _ := strconv.Atoi(block[0])
			StartTime := extractTime(block[1], "START")
			EndTime := extractTime(block[1], "END")
			Content := block[2]
			//put the block in it.
			blocks[len(blocks)-1] = Block{
				Serial:  SNo,
				Start:   StartTime,
				End:     EndTime,
				Content: Content,
			}
		}
	}
	return blocks
}
