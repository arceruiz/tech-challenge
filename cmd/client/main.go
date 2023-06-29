package main

import (
	"client/internal/config"
	"client/internal/rest"

	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	if err := rest.New().Start(); err != nil {
		logrus.Panic()
	}
}
