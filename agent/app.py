import json
import asyncio
from nats.aio.client import Client as NATS
from stan.aio.client import Client as STAN

agent_running = False


async def start_agent():
    cluster_id = "test-cluster"
    client_id = "agent1"
    channel_subject = "agent.1AX24"
    nats_connection = NATS()
    await nats_connection.connect("nats://127.0.0.1:4222")

    sc = STAN()
    await sc.connect(cluster_id, client_id, nats=nats_connection)

    async def cb(msg):
        await sc.ack(msg)
        req_data = json.loads(msg.data.decode())
        global agent_running
        if req_data['event_type'] == 'start' and not agent_running:
            agent_running = True
            config = req_data['config']
            print("Agent started with the following configuration: " + str(config))
            # START AGENT CODE
        elif req_data['event_type'] == 'stop':
            agent_running = False
            print("Agent stopped")
            # STOP AGENT CODE

    await sc.subscribe(channel_subject, manual_acks=True, start_at="last_received", cb=cb)


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(start_agent())
    loop.run_forever()
    loop.close()
