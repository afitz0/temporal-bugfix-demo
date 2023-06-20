package bugfixdemo;

import io.temporal.workflow.WorkflowInterface;
import io.temporal.workflow.WorkflowMethod;

// Workflow interface
@WorkflowInterface
public interface BugFixDemoWorkflow {
    @WorkflowMethod
    String bugfixWorkflow();
}