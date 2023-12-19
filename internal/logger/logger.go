package logger

import (
	"log"
)

type Logger struct {
	Path string
}

func (l Logger) ErrorLog(message string) {
	log.Fatalln(Red + "[ERROR] " + Reset + message)
}

func (l Logger) DNSLog(message string) {
	log.Println(Purple + "[DNS] " + Reset + message)
}

func (l Logger) InfoLog(message string) {
	log.Println(Blue + "[INFO] " + Reset + message)
}

func (l Logger) putInFile(message string) {

}
