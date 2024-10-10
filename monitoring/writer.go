package monitoring

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Writer struct {
	f *os.File
	Input chan string
}

func newWriter(f *os.File) *Writer {
	return &Writer{
		f: f,
		Input: make(chan string),
	}
}

var writer *Writer

//Creates new log file if not present
func InitWriter() *Writer {
	fp, err := os.Getwd()
	if err != nil {
		logrus.Error("Error getting working directory in monitoring, ", err)
	}
	fp = fp + "/monitoring/monitoring.log"
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Error("Cannot open or create file, ", err)
	}
	writer = newWriter(f)
	return writer 
}

func Write(msg string) {
	writer.Input <- msg
}

func (w *Writer) Run() {
	defer w.f.Close()
	for {
		log := <- w.Input
		_, err := w.f.Write([]byte(log))
		if err != nil {
			logrus.Error("Error writing to api log file, ", err)
		}
		w.f.Sync()
	}
}