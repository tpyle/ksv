package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tpyle/ksv/cmd"
	"github.com/tpyle/ksv/lib/cfg"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.TraceLevel)
	cfg, err := cfg.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}
	logrus.Tracef("using config: %+v", cfg)

	cmd.Execute()
}
