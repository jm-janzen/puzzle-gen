package logger

import (
	"io"
	"log"
	"os"
	"time"
)

func InitLogger(fileName string, stdout bool) {
	logFile, err := os.OpenFile(fileName,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0666)
	if err != nil {
		panic(err)
	}
	if stdout {
		log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	} else {
		log.SetOutput(logFile)
	}

	logDatetime()
}

func logDatetime() {
	log.Printf("--- %v ---\n", time.Now().String())
}
