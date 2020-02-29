# Agent-Management

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
