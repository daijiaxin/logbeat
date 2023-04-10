package output

import (
	"logbeat/internal/config"

	"github.com/sirupsen/logrus"
)

var OutputChan = make(map[string]chan map[string]interface{})

func Start(output config.Output) {
	switch output.Type {
	case "file":
		OutputFile(output)
	case "kafka":
		OutputKafka(output)
	case "elasticsearch":
		OutputElasticsearch(output)
	default:
		logrus.WithField("type", output.Type).Warn("output type error, ingored")
	}
}

func StartAll(outputs []config.Output) {
	for _, output := range outputs {
		Start(output)
	}
}
