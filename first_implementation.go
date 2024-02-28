package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type weatherParser struct {
	Name    string
	Weather float64
}

func retrieveValues(m map[string][]float64, err error) {
	if err != nil {
		fmt.Println("Could not retrieve information from file or something else.\nDying...")
	}

	for names, values := range m {
		fmt.Println("-----------------------------------")
		fmt.Println(names, ": ")
		// fmt.Println(values)
		fmt.Println("	min: ", getMinValues(values))
		fmt.Println("	count: ", len(values))
		fmt.Println("	max: ", getMaxValues(values)) //
		fmt.Println("	mean: ", getMean(values))     //
		fmt.Println("-----------------------------------")
	}
}

func getMinValues(values []float64) (min float64) {
	min = 100
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return
}

func getMaxValues(values []float64) (max float64) {
	max = -100
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return
}

func getMean(values []float64) (mean float64) {
	mean = 0
	for _, v := range values {
		mean = mean + v
	}
	mean = (mean / float64(len(values)))
	return
}

func parseFileToMapOrDie(filepath string) (map[string][]float64, error) {
	m := make(map[string][]float64)

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err, "erro abrir file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		records := scanner.Text()
		u := getInfoFromBufferLine(records)

		m[u.Name] = append(m[u.Name], u.Weather)
	}

	return m, nil
}

func convertStringBuffedToFloat(input string) (output float64) {
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Failed to convert string to float")
	}
	return
}

func getInfoFromBufferLine(input string) (data weatherParser) {
	splitedInput := strings.Split(input, ";")
	data.Name = splitedInput[0]
	data.Weather = convertStringBuffedToFloat(splitedInput[1])
	return
}
