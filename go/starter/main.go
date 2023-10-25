package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"starter"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "demo-workflow",
		TaskQueue: "demo",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, starter.Workflow, "Hello", "World")
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// [optionally] Synchronously wait for the workflow completion.
	//	var result string
	//	err = we.Get(context.Background(), &result)
	//	if err != nil {
	//		log.Fatalln("Unable get workflow result", err)
	//	}
	//	log.Println("Workflow result:", result)
}
