package main

import (
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/dig"

	env "./environment"
	msg "./messaging"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	container := buildContainer()

	container.Invoke(func(manager msg.ConsumerManager) {
		manager.Consume("After.debug")
	})
}

func buildContainer() *dig.Container {
	container := dig.New()
	container.Provide(func() env.Manager { return env.NewProduction() })
	container.Provide(func() msg.ConsumerAction { return msg.NewPrintConsumerAction() })
	container.Provide(func(environment env.Manager, action msg.ConsumerAction) msg.ConsumerManager {
		return msg.NewRabbitConsumerManager(environment, action)
	})
	return container
}

func errorHandle(message string, err error) {
	panic(message + ":" + err.Error())
}
