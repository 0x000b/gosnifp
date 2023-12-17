package logger

import "log"

type Logger struct {
	Path string
}

func (l Logger) ErrorLog(message string) {
	log.Println(Red + "[Error] " + Reset + message)
}

func (l Logger) DNSLog(message string) {
	log.Println(Purple + "[DNS] " + Reset + message)
}

func (l Logger) InfoLog(message string) {
	log.Println(Blue + "[Info] " + Reset + message)
}

func (l Logger) putInFile(message string) {

}
