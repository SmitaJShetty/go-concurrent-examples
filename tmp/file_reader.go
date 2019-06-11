package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
)
//FileReader construct for file reader
type FileReader struct {
	logger *log.Logger
}

//NewFileReader construct for filereader
func NewFileReader() *FileReader{
	fr:= &FileReader{
		logger: log.New(),
	}
	fr.logger.SetOutput(os.Stdout)
	fr.logger.SetFormatter(&log.JSONFormatter{})
	return fr
}
//DoWorkChannel reads a file and adds a message to a channel
func (fr *FileReader) DoWorkChannel(fileName string,ch chan int)  {
	fr.logger.WithFields(log.Fields{"info":fileName}).Info("file reading begins...")
	ch <- 1
}

//DoWorkWaitGroup reads a file and adds a message to a channel
func (fr *FileReader) DoWorkWaitGroup(wg *sync.WaitGroup)  {
	fr.logger.WithFields(log.Fields{"info":"no args"}).Info("sthg begins...")
	wg.Done()
}