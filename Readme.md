# Kafka Go Example

This project demonstrates how to send messages to a Kafka topic using Go.

## Prerequisites

- Go installed on your machine
- Docker and Docker Compose installed on your machine

## Running Kafka, Zookeeper, and the Kafka Control Panel with Docker Compose

To start Kafka, Zookeeper, and the Kafka control panel, use the following command:

```sh
docker-compose up -d
```

You can access the control panel at `http://localhost:9021`.

## Running the Go Application

1. Clone the repository:

```sh
git clone https://github.com/joaofilippe/kafka-go.git
cd kafka-go
```

2. Run the Go application:

```sh
go run main.go
```

3. Enter messages in the console to send them to the Kafka topic. You can monitor the messages using the Kafka control panel.

![alt text](/assets/image.png)
![alt text](/assets/image-1.png)