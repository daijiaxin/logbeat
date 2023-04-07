package input

import (
	"io"
	"logbeat/internal/config"
	"logbeat/internal/filter"
	"time"

	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
)

var pathChanMap = make(map[string]chan int)

func InputFile(input config.Input) {
	for _, path := range input.Path {
		if _, ok := pathChanMap[path]; !ok {
			quit := make(chan int)
			go InputFileWorker(path, input.Codec, quit)
			pathChanMap[path] = quit
		}
	}
}

func InputFileWorker(path string, codec string, quit <-chan int) {
	fields := logrus.Fields{
		"type": "file",
		"path": path,
	}
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: io.SeekEnd,
		},
		MustExist: false,
		Poll:      true,
	}

	tails, err := tail.TailFile(path, config)
	if err != nil {
		logrus.WithFields(fields).Error("input goroutine error")
	}

	var line *tail.Line
	var ok bool
	for {
		select {
		case <-quit:
			logrus.WithFields(fields).Info("input goroutine quit")
			return
		default:
			line, ok = <-tails.Lines
			if !ok {
				time.Sleep(time.Second)
				continue
			}
			filter.FilterChan <- Codec([]byte(line.Text), codec)
		}
	}
}
