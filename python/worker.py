import asyncio
import logging
from datetime import timedelta

from workflow import *

from temporalio.client import Client
from temporalio.worker import Worker


async def main():
    logging.basicConfig(level=logging.INFO)

    # Start client
    client = await Client.connect("localhost:7233")

    worker = Worker(
        client,
        task_queue="bugfix-task-queue",
        workflows=[BugfixWorkflow],
        activities=[one, two, three],
        sticky_queue_schedule_to_start_timeout=timedelta(seconds=1),
    )
    logging.info("Starting worker.")
    await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
