package main

import (
	"bytes"
	"fmt"
	"io"
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
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("%s submitted %d %s\n", sendTime, hws.Id, hws.Comment))
}

func (hwa *HwAccepted) Send() []byte {
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	return []byte(fmt.Sprintf("%s accepted %d %d\n", sendTime, hwa.Id, hwa.Grade))
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
