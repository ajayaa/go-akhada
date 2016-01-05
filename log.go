package main

import (
	"bytes"
	//"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello world")
	f, err := os.Create("/tmp/test")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.Write(buf.Bytes())
	i := 0
	for {
		if i == 10 {
			break
		}
		f.Write([]byte(time.Now().String() + "\n"))
		i++
	}
}
