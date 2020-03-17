![](https://raw.githubusercontent.com/nats-io/nats-site/master/src/img/large-logo.png)

# Agent-Management
NATS is a simple, secure and high performance open source messaging system for cloud native applications, IoT messaging, and micro-services architectures.Whereas NATS Streaming is a data streaming system powered by NATS, and written in the Go programming language.For more detail chekout this [link](https://medium.com/@syedrizwan161/high-performance-messaging-with-nats-1c261fae3778)

## Run NATS streaming server using the following command. 

`docker run -p 4222:4222 -p 8222:8222 nats-streaming -store file -m 8222 -dir datastore`

## Nats Streaming Pub-Sub Pattern

1. Start subscriber/agent 
    ```
    cd agent
    pip install -r requiremnts.txt
    python app.py
    ```

2. Run Publisher/manager
    ```
    cd ..
    cd manager/publisher
    go run publisher.go
    ```

## Nats Request-Reply Pattern
1. Start agent 
    ```
    cd agent
    pip install -r requiremnts.txt
    python app.py
    ```

2. Run manager
    ```
    cd ..
    cd manager/request
    go run request.go
    ```
