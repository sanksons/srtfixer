package srt

import (

	"io/ioutil"
	"strings"
	//"srtfixer/utility"
	//"fmt"
	"strconv"
	"time"
	"os"
	//"fmt"

	//"path/filepath"
	"fmt"
)


type Block struct {
	Serial  int
	Start   time.Time
	End     time.Time
	Content string
}

//Fetch the contents of the specified file.
func FetchFileContents(filePath string) (string, error) {
	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func WriteData(dataArr []Block, filepath string) bool{

	csvfile, er := os.OpenFile(filepath, os.O_RDWR, 0777)
	fmt.Printf("%v", er)
	dataStr := ""
	for _,e := range dataArr {
		dataStr = dataStr + strconv.Itoa(e.Serial) + "\r\n"
		milliSecString := strconv.Itoa(e.Start.Nanosecond()/1000000)
		dataStr = dataStr + e.Start.Format("15:04:05") + "," + milliSecString
				+ " --> " + e.End.Format("15:04:05") + "\r\n"
		dataStr = dataStr + e.Content + "\r\n\r\n"
	}
	_ , er = csvfile.Write([]byte(dataStr))
	fmt.Printf("%v", er)
	defer csvfile.Close()
	return true;
}

//Parse Supplied data
//@todo: work on datetime go functions
func ParseData(content string) []Block {
	//remove carriage return chars first
	content = strings.Replace(content, "\r", "", -1)
	//Split into blocks containing three sets
	contentBlocks := strings.Split(content, "\n\n")

	//declare a Slice to store Blocks.
	var mainC []Block
	for _, e := range contentBlocks {
		block := strings.SplitN(e, "\n", 3)
        if len(block) >= 3 {
			serial, _ := strconv.Atoi(block[0])
			dateArr := strings.Split(block[1], " --> ")
			//Extend our container array
			mainCNew := make([]Block, len(mainC)+1)
			for i := range mainC {
				mainCNew[i] = mainC[i]
			}
			mainC = mainCNew

			startDate := strings.Split(dateArr[0],",")
			startDateTime,_ := time.Parse("15:04:05", startDate[0])

			//add nanoseconds to thi s tim
			startDateTime.Nanosecond()

			intMicroSeconds,_ := strconv.Atoi(startDate[1])
			startDateTime = startDateTime.Add(time.Duration(intMicroSeconds) * 1000000)

			endDate := strings.Split(dateArr[1],",")
			endDateTime,_ := time.Parse("15:04:05", endDate[0])
			intMicroSeconds,_ = strconv.Atoi(endDate[1])
			endDateTime = endDateTime.Add(time.Duration(intMicroSeconds) * 1000000)

			//put the block in it.
			mainC[len(mainC)-1] = Block{
				Serial  : serial,
				Start   : startDateTime,
				End     : endDateTime,
				Content : block[2],
			}
		}
		//fmt.Printf("%v", mainC)
	}
	return mainC
}

//prepare data based on lag
func AddLagToTime(originalData []Block, lag time.Duration ) []Block {

	for i,e := range originalData {
		originalData[i].Start = e.Start.Add(lag)
		originalData[i].End   = e.End.Add(lag)
	}
	return originalData
}



