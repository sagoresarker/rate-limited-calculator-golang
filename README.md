# Rate-limited Calculator Doccumentation

## Table of Contents

- [Project Details](#project-details)
  - [About the Project](#about-the-project)
  - [How Sliding Window Algorithm Works](#how-sliding-window-algorithm-works)
  - [Architecture of the Implemented System](#architecture-of-the-implemented-system)
- [Run and Test the Project](#run-and-test-the-project)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Running Unit Tests](#running-unit-tests)
  - [Endpoints](#endpoints)

Before diving into setup and test, firt describe a bit more about this project.

## What I exactly implement here
I implemented a rate-limited calculator service in GoLang, utilizing a sliding window algorithm for rate limiting and Redis for data storage.

## How Sliding Window Algorithm Works Here
In this system, the sliding window algorithm works by maintaining a record of request timestamps for each user within a specified time window (e.g., one minute). When a new request comes in, the algorithm checks the timestamps and removes any that fall outside the current window. It then counts the remaining timestamps. If the count is below the allowed limit, the request is accepted and the current timestamp is added to the record. If the count exceeds the limit, the request is denied.
## Archetecture of the implemented system
In my system, the sliding window algorithm is used to limit the number of requests a user can make within a given time frame. Here's how it works:

![Sliding Window](docs/img/ratelimiter-sliding-window-algo.png)




## Prerequisite
1. redis
2. docker
3. Go
4. make (not required)


## Setup

To get started with this project, follow these steps:

1. Clone this repository to local machine using `git clone https://github.com/sagoresarker/rate-limited-calculator-golang.git`.
2. Navigate to the project directory: `cd rate-limited-calculator-golang`.
3. Install project dependencies by running: `go mod tidy`. (sometime it required to remove ```go.sum``` file before run this)
4. Start redis container: `docker-compose up -d`

## Run unit test
To run unit test, if you already have `make` tool on your system, you can run the test by just

```bash
make test
```

If you don't have make in your system, just run
```bash
go test ./...
```

## Endpoints

This document provides a brief overview of the one API endpoints in the `handlers.go` file.

### Payload structure
Payload contain the operation name(calculator_type), username, and numbers.

```json
{
    "username": "testuser",
    "calculator_type": "add",
    "number1": 10,
    "number2": 5
}
```

Currently there are ```7``` calculator_type, these are
```bash
calculator_type
- add
- subtract
- multiply
- divide
- modulo
- power
- factorial
```

### 1. POST /calculate

This endpoint is used to send a ```POST``` request to ```/calculator``` endpoint with a specific configuration.

**Request Body:** This endpoint require ```username, calculator_type, numbers``` in the  payload in the POST request.

**Response:** The response will contain the status of the operation and cluster name and userid

Here is the `curl` commands to test this  endpoint locally for addition:

### Addition

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "add",
    "number1": 10,
    "number2": 5
}'
```

### Subtraction

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "subtract",
    "number1": 10,
    "number2": 5
}'
```

### Multiplication

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "multiply",
    "number1": 10,
    "number2": 5
}'
```


### Division

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "divide",
    "number1": 10,
    "number2": 5
}'
```

### Modulo

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "modulo",
    "number1": 10,
    "number2": 5
}'
```

### Power

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "power",
    "number1": 10,
    "number2": 5
}'
```

### Factorial

```bash
curl -X POST http://localhost:8080/calculate \
-H "Content-Type: application/json" \
-d '{
    "username": "testuser",
    "calculator_type": "factorial",
    "number1": 10,
    "number2": 5
}'
```


When it reaches it limits, the server will return a ```429 Too Many Request```


Thats it






