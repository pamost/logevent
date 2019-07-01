package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
)

type Event interface {
	Send() []byte
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type HwAccepted struct {
	Id    int
	Grade int
}

func (hws *HwSubmitted) Send() []byte {
	sendTime := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	id := strconv.Itoa(hws.Id)
	comment := string(hws.Comment)
	return []byte(fmt.Sprintf("%s submitted %s %s\n", sendTime, id, comment))
}

func (hwa *HwAccepted) Send() []byte {
	sendTime := time.Now().Format("2006-01-02")
	id := strconv.Itoa(hwa.Id)
	grade := strconv.Itoa(hwa.Grade)
	return []byte(fmt.Sprintf("%s accepted %s %s\n", sendTime, id, grade))
}

func LogEvent(e Event, w io.Writer) {
	w.Write(e.Send())
}

func main() {

	var bufferSubmit, bufferAccept bytes.Buffer

	//2019-01-01 submitted 3456 "please take a look at my homework"
	dataSubmit := &HwSubmitted{3456, "ABCD", "please take a look at my homework"}
	LogEvent(dataSubmit, &bufferSubmit)
	fmt.Printf("%s", bufferSubmit.String())

	// 2019-01-01 accepted 3456 4
	dataAccept := &HwAccepted{3456, 4}
	LogEvent(dataAccept, &bufferAccept)
	fmt.Printf("%s", bufferAccept.String())
}
