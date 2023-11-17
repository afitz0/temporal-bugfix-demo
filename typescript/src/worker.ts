import { DefaultLogger, NativeConnection, Worker, Runtime } from '@temporalio/worker';
import * as activities from './activities';

async function run() {
  Runtime.install({ logger: new DefaultLogger('WARN') });

  const connection = await NativeConnection.connect({
    address: 'localhost:7233',
  });

  const worker = await Worker.create({
    connection,
    namespace: 'default',
    taskQueue: 'bugfix-demo',
    workflowsPath: require.resolve('./workflows'),
    stickyQueueScheduleToStartTimeout: '1s',
    activities,
  });

  await worker.run();
}

run().catch((err) => {
  console.error(err);
  process.exit(1);
});
