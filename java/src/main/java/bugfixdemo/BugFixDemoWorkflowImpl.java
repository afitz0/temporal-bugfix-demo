package bugfixdemo;

import io.temporal.activity.ActivityOptions;
import io.temporal.common.RetryOptions;
import io.temporal.workflow.Workflow;

import java.time.Duration;

public class BugFixDemoWorkflowImpl implements BugFixDemoWorkflow {

    private final BugFixDemoActivities activities =
            Workflow.newActivityStub(
                    BugFixDemoActivities.class,
                    ActivityOptions.newBuilder()
                            .setStartToCloseTimeout(Duration.ofSeconds(2))
                            .setRetryOptions(RetryOptions.newBuilder()
                                    .setInitialInterval(Duration.ofSeconds(1))
                                    .setBackoffCoefficient(1.0)
                                    .build())
                            .build());

    @Override
    public String bugfixWorkflow() {
        activities.stepOne("First one!");
        activities.stepTwo("Second one!");
        activities.stepThree("Last one!");
        return "Success";
    }
}
