package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var Log *log.Logger

func InitLogger() {

	if err := os.Mkdir("Logs-go-gin", os.ModePerm); err != nil {
		fmt.Println(err)
	}

	logFilePath := filepath.Join("./Logs-go-gin", "text.log")
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error occur in file opening of text.log file", err.Error())
	}
	log.SetOutput(f)
	Log = log.New(f, "", log.LstdFlags)
}
