import asyncio

from datetime import timedelta

from temporalio import activity, workflow
from temporalio.common import RetryPolicy


@workflow.defn
class BugfixWorkflow:
    @workflow.run
    async def run(self) -> None:
        workflow.logger.info("Running workflow")

        rt_policy = RetryPolicy(
            maximum_interval=timedelta(seconds=1),
            backoff_coefficient=1.0)

        await workflow.execute_activity(
            one,
            start_to_close_timeout=timedelta(seconds=1),
            retry_policy=rt_policy,
        )

        await workflow.execute_activity(
            two,
            start_to_close_timeout=timedelta(seconds=1),
            heartbeat_timeout=timedelta(seconds=1),
            retry_policy=rt_policy,
        )

        await workflow.execute_activity(
            three,
            start_to_close_timeout=timedelta(seconds=1),
            retry_policy=rt_policy,
        )


@activity.defn
async def one() -> str:
    print("--------> Running one!")
    return "One done!"


@activity.defn
async def two() -> str:
    print("--------> Running two!")

    has_bug = True
    if has_bug:
        raise TypeError
    else:
        return "Two done!"


@activity.defn
async def three() -> str:
    print("--------> Running three!")
    return "Three done!"
