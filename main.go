package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// QueueService represents a simple queue service implemented using goroutines and channels.
type QueueService struct {
	inputChannel  chan string
	outputChannel chan string
}

// NewQueueService creates a new instance of the QueueService.
func NewQueueService() *QueueService {
	return &QueueService{
		inputChannel:  make(chan string),
		outputChannel: make(chan string),
	}
}

// Start starts the queue service.
func (qs *QueueService) Start() {
	go func() {
		for {
			message := <-qs.inputChannel
			// Simulate some processing time
			time.Sleep(time.Millisecond * 500)
			qs.outputChannel <- message
		}
	}()
}

// Enqueue sends a message to the queue.
func (qs *QueueService) Enqueue(message string) {
	qs.inputChannel <- message
}

// Dequeue receives a message from the queue.
func (qs *QueueService) Dequeue() string {
	return <-qs.outputChannel
}

func main() {
	// Check if the correct number of command-line arguments is provided
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		os.Exit(1)
	}

	// Create a new queue service
	queueService := NewQueueService()
	queueService.Start()

	// Worker 1: Read lines from a file and write to the queue service
	go func() {
		inputFile, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("Error opening input file:", err)
			return
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)
		for scanner.Scan() {
			line := scanner.Text()
			queueService.Enqueue(line)
			fmt.Println(line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input file:", err)
		}
	}()

	// Worker 2: Read lines from the queue service and write to a file
	go func() {
		outputFile, err := os.Create(os.Args[2])
		if err != nil {
			fmt.Println("Error creating output file:", err)
			return
		}
		defer outputFile.Close()

		for {
			line := queueService.Dequeue()
			_, err := outputFile.WriteString(line + "\n")
			if err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
		}
	}()

	// Sleep to allow workers to finish processing
	time.Sleep(time.Second * 5)
}
