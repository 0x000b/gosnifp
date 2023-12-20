package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

type Logger struct {
	Path string
}

func (l Logger) ErrorLog(message string, err string) {
	log.Error(message, "err", err)
}

func (l Logger) FatalLog(message string, err error) {
	log.Fatal(message, "fatal", err)
}

func (l Logger) DNSLog(message string) {
	log.Print(message)
	l.putInFile(message)
}

func (l Logger) InfoLog(message string) {
	log.Info(message)
}

func (l Logger) putInFile(message string) {
	time := time.Now().Format(time.DateOnly)
	f, err := os.OpenFile(l.Path+"/"+time+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		l.ErrorLog("Error in file creation: ", err.Error())
	}

	defer f.Close()

	if _, err := f.WriteString(message + "\n"); err != nil {
		l.ErrorLog("Error while writing to the file: ", err.Error())
	}
}
