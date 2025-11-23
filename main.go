package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tpyle/ksv/cmd"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.TraceLevel)

	cmd.Execute()
}
