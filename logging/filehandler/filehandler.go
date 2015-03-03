package filehandler

import (
	"github.com/henkburgstra/go/logging"
	"os"
)

type FileHandler struct {
	logging.StreamHandler
	filename string
	mode     string
}

func NewFileHandler(filename string, mode string) *FileHandler {
	fileHandler := new(FileHandler)
	fileHandler.filename = filename
	fileHandler.mode = mode

	logfile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0666))
	fileHandler.StreamHandler = *logging.NewStreamHandler(logfile)

	return fileHandler
}
