package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tpyle/ksv/cmd"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.TraceLevel)

	cmd.Execute()
}
