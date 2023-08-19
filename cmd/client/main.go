package main

import (
	"fmt"
	"tech-challenge/internal/config"
	"tech-challenge/internal/rest"

	"github.com/sirupsen/logrus"
)

func main() {
	config.ParseFromFlags()

	if err := rest.New().Start(); err != nil {
		logrus.Panic()
	}
}

func testefunc() {
	fmt.Printf("Olá Rapazes gostosos")
}
