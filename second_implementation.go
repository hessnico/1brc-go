package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type weatherParsingInfo struct {
	max   float64
	min   float64
	sum   float64
	count float64
}

func roundOneDecimal(input float64) (output float64) {
	output = math.Round(input*10) / 10
	return
}

func retrieveValues(mapTemp map[string]weatherParsingInfo, err error) {
	if err != nil {
		fmt.Println("Could not retrieve information from file or something else.\nDying...")
	}

	for names, values := range mapTemp {
		fmt.Println("-----------------------------------")
		fmt.Println(names, ": ")
		fmt.Println("	max: ", values.max)
		fmt.Println("	min: ", values.min)
		fmt.Println("	avg: ", roundOneDecimal(values.sum/values.count))
		fmt.Println("	count: ", values.count)
		fmt.Println("-----------------------------------")
	}
}

func parseFileToMapOrDie(filepath string) (map[string]weatherParsingInfo, error) {

	mapTemp := make(map[string]weatherParsingInfo)

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err, "erro abrir file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		records := scanner.Text()
		splitedRecords := strings.Split(records, ";")

		weather_name := splitedRecords[0]
		weather_temp := convertStringBuffedToFloat(splitedRecords[1])

		if value, ok := mapTemp[weather_name]; ok {
			if weather_temp > value.max {
				value.max = weather_temp
			}

			if weather_temp < value.min {
				value.min = weather_temp
			}

			value.sum += weather_temp
			value.count++

			mapTemp[weather_name] = value
		} else {
			mapTemp[weather_name] = weatherParsingInfo{
				max:   weather_temp,
				min:   weather_temp,
				sum:   weather_temp,
				count: 1,
			}
		}
	}

	return mapTemp, nil
}

func convertStringBuffedToFloat(input string) (output float64) {
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Failed to convert string to float")
	}
	return
}
