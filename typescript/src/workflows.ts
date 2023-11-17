import * as wf from '@temporalio/workflow';
import type * as activities from './activities';

const { stepOne, stepTwo, stepThree } = wf.proxyActivities<typeof activities>({
  startToCloseTimeout: '1s',
  retry: {
      backoffCoefficient: 1,
      initialInterval: '1s',
  },
});

export async function bugFixWorkflow(): Promise<string> {
  let result = await stepOne("1");
  wf.log.info(`Result from step one: ${result}`)
 
  result = await stepTwo("2");
  wf.log.info(`Result from step two: ${result}`)

  result = await stepThree("3");
  wf.log.info(`Result from step three: ${result}`)

  return "Done!"
}
