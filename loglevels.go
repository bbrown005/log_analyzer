package main

import "strings"

func CountLogLevels(lines []string) map[string]int {
	counts := make(map[string]int)
	for _, line := range lines {
		words := strings.Fields(line)
		level := words[0]
		counts[level]++
	}
	return counts
}
