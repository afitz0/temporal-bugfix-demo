package bugfixdemo;

import io.temporal.api.common.v1.WorkflowExecution;
import io.temporal.client.WorkflowClient;
import io.temporal.client.WorkflowOptions;
import io.temporal.serviceclient.WorkflowServiceStubs;
import io.temporal.worker.Worker;
import io.temporal.worker.WorkerFactory;

public class Starter {
    static final String WORKFLOW_ID = "bugfixdemo-workflow";

    public static void main(String[] args) {
        WorkflowServiceStubs service = WorkflowServiceStubs.newLocalServiceStubs();
        WorkflowClient client = WorkflowClient.newInstance(service);

        WorkflowOptions workflowOptions =
                WorkflowOptions.newBuilder().setWorkflowId(WORKFLOW_ID).setTaskQueue(Shared.TASK_QUEUE_NAME).build();
        BugFixDemoWorkflow workflow = client.newWorkflowStub(BugFixDemoWorkflow.class, workflowOptions);

        WorkflowExecution we =  WorkflowClient.start(workflow::bugfixWorkflow);
        System.out.println("Started the workflow. See the Worker's console for output.");
        System.out.printf("Workflow ID: %s\nRun ID: %s\n", we.getWorkflowId(), we.getRunId());

        System.exit(0);
    }
}