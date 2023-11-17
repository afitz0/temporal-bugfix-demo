import { Connection, Client } from '@temporalio/client';
import { bugFixWorkflow } from './workflows';
import { nanoid } from 'nanoid';

async function run() {
  const connection = await Connection.connect({ address: 'localhost:7233' });

  const client = new Client({
    connection,
  });

  const handle = await client.workflow.start(bugFixWorkflow, {
    taskQueue: 'bugfix-demo',
    args: [],
    workflowId: 'workflow-' + nanoid(),
  });
  console.log(`Started workflow ${handle.workflowId}.`);
  console.log(`Visit http://localhost:8233/namespaces/default/workflows/${handle.workflowId}`);

  // optional: wait for client result
  //console.log(await handle.result());
}

run().catch((err) => {
  console.error(err);
  process.exit(1);
});
