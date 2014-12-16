package main

import (
	"fmt"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	//"io"
	//"io/ioutil"
	//"log"
	"os"
	//"regexp"
	"strings"
	"srtfixer/srt"
	//"time"
)

type Person struct {
	Name  string
	Voila int
}

type Config struct {
	File       string
	Start      string
	End        string
	Lag        string
	IsPositive string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func prepareArgs(args []string) Config {

	configStruct := Config{}

	for _, element := range args {
		splitElement := strings.Split(element, "=")
		//fmt.Println(splitElement)

		if strings.Contains(splitElement[0], "--file") {
			configStruct.File = splitElement[1]
		} else if strings.Contains(splitElement[0], "--start") {
			configStruct.Start = splitElement[1]
		} else if strings.Contains(splitElement[0], "--end") {
			configStruct.End = splitElement[1]
		} else {
			//what is this.
		}
	}

	return configStruct
}

func (c Config) isValid() bool {

	returnVar := true
	//check if the supplied config is valid.

	//check if supplied file is valid
	if _, err := os.Stat(c.File); err != nil {
		returnVar = false
	}
	return returnVar
}

func main() {


	da,_ := srt.FetchFileContents("/home/jackson/test2.srt")
	de := srt.ParseData(da)
	du:= srt.AddLagToTime(de, -1000000000)
	srt.WriteData(du, "/home/jackson/test.srt")
	fmt.Println("")
	return
/*
	csvfile, err := os.Open("/home/jackson/test.srt")


	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()



	dat, err := ioutil.ReadFile("/home/jackson/test.srt")

	dat1 := strings.Replace(string(dat), "\r", "", -1)
	//contents := RegSplit(dat1, "\n\n")
	for _, element := range contents {
		fmt.Println(element)
		fmt.Println("sdsdsdsdsdsdds")
	}
	return
	fmt.Printf("%+v", contents)
	return

	cc := prepareArgs(os.Args).isValid()
	fmt.Printf("%+v", cc)
	return
	for index, element := range os.Args {
		fmt.Println(index)
		fmt.Println(element)
	}
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]
	//a := os.Args[1]
	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)
	//fmt.Println(a)

	return

	session, err := mgo.Dial("localhost")

	if err != nil {
		fmt.Println("Connection failed")
	}

	result := Person{}
	testData := session.DB("mydb").C("testData")
	testData.Insert()
	err1 := session.DB("mydb").C("testData").Find(bson.M{}).One(&result)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Printf("%+v", result)
*/
}
