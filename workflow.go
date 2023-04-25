package starter

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context, greeting string, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second,
			BackoffCoefficient: 1.0,
			MaximumInterval:    10 * time.Second,
			MaximumAttempts:    0, // 0 is infinite
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Starter workflow started", "greeting", greeting, "name", name)

	var a *Activities

	var result string
	err := workflow.ExecuteActivity(ctx, a.NormalActivity, "First", "Time").Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	fmt.Println("Result from first activity:", result)

	err = workflow.ExecuteActivity(ctx, a.BuggyActivity).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	fmt.Println("Result from second activity:", result)

	err = workflow.ExecuteActivity(ctx, a.NormalActivity, "Second", "Time").Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}

	fmt.Println("Result from third activity:", result)

	logger.Info("Starter workflow completed.", "result", result)
	return result, nil
}
