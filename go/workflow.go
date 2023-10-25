package starter

import (
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
			MaximumInterval:    time.Second,
			MaximumAttempts:    0, // 0 is infinite
		},
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Starter workflow started", "greeting", greeting, "name", name)

	var a *Activities

	var result string

	// ------------- Step 1 ------------- //
	err := workflow.ExecuteActivity(ctx,
		a.StepOne, "First!").Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("[Workflow] Result from first step:", result)

	// ------------- Step 2 ------------- //
	err = workflow.ExecuteActivity(ctx,
		a.StepTwo, "Second!").Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("[Workflow] Result from second step:", result)

	// ------------- Step 3 ------------- //
	err = workflow.ExecuteActivity(ctx,
		a.StepThree, "Third!").Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("[Workflow] Result from third step:", result)

	logger.Info("Starter workflow completed.", "result", result)
	return result, nil
}
