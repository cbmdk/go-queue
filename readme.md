# Simple Queue service

## Task Objective
Read lines from a file

Write lines to a queue service via a network protocol

Read lines from the queue service via a network protocol

Write lines to a file

Implement it as 2 asynchronous workers exchanging information by using a service.

Queuing service needs to be written with only stdlib of Go

## Go Installation

If on a Mac, simply run 

```bash
brew install go
```
Make sure that your Go version is >= 1.21.5 by running 
```bash
go version
```


## Usage

To execute the code  run 
```bash
go run main.go input.txt output.txt
```

To test the code run
```bash
go test
```
