import asyncio
import logging
from temporalio.client import Client

from workflow import *


async def main() -> None:
    logging.basicConfig(level=logging.INFO)

    client = await Client.connect("localhost:7233")

    await client.start_workflow(
        BugfixWorkflow.run,
        id="bugfix-workflow",
        task_queue="bugfix-task-queue",
    )


if __name__ == "__main__":
    asyncio.run(main())
