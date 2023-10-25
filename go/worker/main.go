package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/zap/zapcore"

	"starter"
	"starter/zapadapter"
)

func main() {
	c, err := client.NewLazyClient(client.Options{
		Logger: zapadapter.NewZapAdapter(
			zapadapter.NewZapLogger(zapcore.FatalLevel)),
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "demo", worker.Options{})

	a := &starter.Activities{}
	w.RegisterWorkflow(starter.Workflow)
	w.RegisterActivity(a)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
