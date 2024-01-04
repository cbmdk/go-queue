package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestQueueService(t *testing.T) {
	// Create a new queue service
	queueService := NewQueueService()
	queueService.Start()

	// Worker 1: Enqueue messages
	go func() {
		for i := 1; i <= 3; i++ {
			queueService.Enqueue(fmt.Sprintf("Message %d", i))
		}
	}()

	// Worker 2: Dequeue messages
	go func() {
		for i := 1; i <= 3; i++ {
			message := queueService.Dequeue()
			expectedMessage := fmt.Sprintf("Message %d", i)
			if message != expectedMessage {
				t.Errorf("Expected %s, but got %s", expectedMessage, message)
			}
		}
	}()

	// Sleep to allow workers to finish processing
	time.Sleep(time.Second * 2)
}

func TestFileReadWrite(t *testing.T) {
	// Create a temporary input file
	inputFileName := "test_input.txt"
	defer os.Remove(inputFileName)

	// Write test data to the input file
	inputFile, err := os.Create(inputFileName)
	if err != nil {
		t.Fatal("Error creating test input file:", err)
	}
	defer inputFile.Close()
	inputFile.WriteString("Line 1\nLine 2\nLine 3\n")

	// Create a temporary output file
	outputFileName := "test_output.txt"
	defer os.Remove(outputFileName)

	// Run the main logic with the test files
	go func() {
		os.Args = []string{"", inputFileName, outputFileName}
		main()
	}()

	// Sleep to allow the main logic to finish processing
	time.Sleep(time.Second * 2)

	// Read the output file and check the content
	outputFile, err := os.Open(outputFileName)
	if err != nil {
		t.Fatal("Error opening test output file:", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(outputFile)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	for i, expected := range expectedLines {
		if i >= len(lines) {
			t.Errorf("Expected %s, but no more lines in the output file", expected)
			continue
		}
		if lines[i] != expected {
			t.Errorf("Expected %s, but got %s", expected, lines[i])
		}
	}
}
