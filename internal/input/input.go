package input

import (
	"logbeat/internal/config"

	"github.com/sirupsen/logrus"
)

func Start(input config.Input) {
	switch input.Type {
	case "file":
		InputFile(input)
	case "kafka":
		InputKafka(input)
	default:
		logrus.WithField("type", input.Type).Warn("input type error, ingored")
	}
}

func StartAll(inputs []config.Input) {
	for _, input := range inputs {
		Start(input)
	}
}
