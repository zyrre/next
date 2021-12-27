package utils

import (
	"bufio"
	"os"
	"strings"
)

func FileToMap(filename string) (map[string][]string, error) {
	lines, err := parseFile(filename)
	if err != nil {
		return nil, err
	}
	return mdDataToMap(lines), nil
}

func mdDataToMap(mdData []string) map[string][]string {
	mdMap := make(map[string][]string)
	var currentIndex string
	for _, line := range mdData {
		splitLine := strings.SplitN(line, " ", 2)
		if splitLine[0] == "#" {
			mdMap[splitLine[1]] = []string{}
			currentIndex = splitLine[1]
		} else {
			mdMap[currentIndex] = append(mdMap[currentIndex], strings.Trim(splitLine[1], "~"))
		}
	}
	return mdMap
}

//parse a file and return each non-whitespace line
func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}
