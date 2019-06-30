package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"time"
)

var (
	bufferSubmit, bufferAccept bytes.Buffer
	sendTime                   string = time.Now().Format("2006-01-02")
	submitted                  string = "submitted"
	accepted                   string = "accepted"
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
	id := strconv.Itoa(hws.Id)
	comment := string(hws.Comment)
	return []byte(sendTime + " " + submitted + " " + id + " " + comment + "\n")
}

func (hwa *HwAccepted) Send() []byte {
	id := strconv.Itoa(hwa.Id)
	grade := strconv.Itoa(hwa.Id)
	return []byte(sendTime + " " + accepted + " " + id + " " + grade + "\n")
}

func LogEvent(e Event, w io.Writer) {
	w.Write(e.Send())
}

func main() {
	//2019-01-01 submitted 3456 "please take a look at my homework"
	dataSubmit := &HwSubmitted{3456, "ABCD", "please take a look at my homework"}
	LogEvent(dataSubmit, &bufferSubmit)
	fmt.Printf("%s", bufferSubmit.String())

	// 2019-01-01 accepted 3456 4
	dataAccept := &HwAccepted{3456, 4}
	LogEvent(dataAccept, &bufferAccept)
	fmt.Printf("%s", bufferAccept.String())
}
