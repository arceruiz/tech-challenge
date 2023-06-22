package main

import (
	"client/internal/rest"

	"github.com/sirupsen/logrus"
)

func main() {
	if err := rest.New().Start(); err != nil {
		logrus.Panic()
	}
}
