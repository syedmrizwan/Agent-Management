import asyncio
from nats.aio.client import Client as NATS


async def get_health():
    channel_subject = "agent.1AX24.health"
    nats_connection = NATS()
    await nats_connection.connect("nats://127.0.0.1:4222")

    async def help_request(msg):
        print(msg.data)
        reply = msg.reply
        await nats_connection.publish(reply, b'Agent.1AX24 is healthy.')

    await nats_connection.subscribe(channel_subject, "workers", help_request)

    print("Listening for requests on '{0}' subject...".format(channel_subject))


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(get_health())
    loop.run_forever()
    loop.close()
