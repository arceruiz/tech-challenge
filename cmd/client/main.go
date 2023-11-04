package main

import (
	"os"
	"os/signal"
	"tech-challenge/internal/channels/grpc"
	"tech-challenge/internal/channels/rest"
	"tech-challenge/internal/config"

	"github.com/sirupsen/logrus"
)

var (
	cfg = &config.Cfg
)

func main() {
	config.ParseFromFlags()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := rest.New().Start(); err != nil {
			logrus.Panic()
		}
	}()

	go func() {
		if err := grpc.Listen(cfg.Server.CustomerPort); err != nil {
			logrus.Panic()
		}
	}()

	logrus.WithField("grpc server started on: ", cfg.Server.CustomerPort).Info()
	<-stop
}
