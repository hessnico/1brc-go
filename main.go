package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// <count>/<min>/<mean>/<max>/

func main() {
	logger := log.New(os.Stdout, "", 1)
	f, err := os.OpenFile("./time_logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("error opening save file: %v", err)
	}
	defer f.Close()

	wrt := io.MultiWriter(os.Stdout, f)
	logger.SetOutput(wrt)

	filepaths := [][]string{
		//		{"./data/measurements10.txt", "10"},
		//		{"./data/measurements100.txt", "100"},
		//		{"./data/measurements1000.txt", "1000"},
		//		{"./data/measurements10000.txt", "10000"},
		//		{"./data/measurements100000.txt", "100000"},
		//		{"./data/measurements1000000.txt", "1000000"},
		//		{"./data/measurements10000000.txt", "10000000"},
		//		{"./data/measurements100000000.txt", "10000000"},
		{"./data/measurements1000000000.txt", "100000000"},
	}

	logger.Println("\n")
	for _, file := range filepaths {
		calculate_time_per_run(file[0], file[1], logger)
	}
}

func calculate_time_per_run(filename string, n_rum string, logger *log.Logger) {
	start := time.Now()

	first_impl_run(filename, logger)

	t := time.Now()
	elapsed := t.Sub(start)
	logger.Println("				Time elapsed:", n_rum, ":", elapsed)
}

func first_impl_run(filepath string, logger *log.Logger) {
	map_, err := parseFileToMapOrDie(filepath)
	if err != nil {
		fmt.Println("fail to parse file into map")
	}
	retrieveValues(map_, err)
	map_ = nil
}
