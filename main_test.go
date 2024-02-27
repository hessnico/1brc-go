package main

import (
	"testing"
)

func BenchmarParseFileToMapOrDie(b *testing.B) {
	var datasets [1]string
	datasets[0] = "./data/measurements10.txt"

	for _, d := range datasets {
		parseFileToMapOrDie(d)
	}
}
