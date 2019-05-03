package libs

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	Info    string
	Error   string
	Warning string
}

func NewLogger() *Logger {
	return &Logger{
		Info:    "info",
		Error:   "error",
		Warning: "warning",
	}
}

func (l Logger) I(s string) {
	f, err := os.OpenFile("./logs/"+l.Info, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
}

func (l Logger) E(s string) {
	f, err := os.OpenFile("./logs/"+l.Error, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	// log.Println("Server v%s pid=%d started with processes: %d", "0.0.1", os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
}

func (l Logger) W(s string) {
	f, err := os.OpenFile("./logs/"+l.Warning, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
}
