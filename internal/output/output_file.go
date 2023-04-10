package output

import (
	"encoding/json"
	"fmt"
	"logbeat/internal/config"
	"os"
	"time"
)

func OutputFile(output config.Output) {
	c := make(chan map[string]interface{}, 10000)
	key := fmt.Sprintf("%s-%s-%s", output.Type, output.Codec, output.Path)
	OutputChan[key] = c
	go OutputFileWork(output, c)
}

func OutputFileWork(output config.Output, c chan map[string]interface{}) {
	file := make(map[string]*os.File)
	for {
		select {
		case msg := <-c:
			b, err := json.Marshal(msg)
			if err != nil {
				continue
			}
			var f *os.File
			if _, ok := file[output.Path]; ok {
				f = file[output.Path]
			} else {
				f, _ = os.OpenFile(output.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
				file[output.Path] = f
			}
			f.Write(b)
		default:
			time.Sleep(time.Second)
		}
	}
}
