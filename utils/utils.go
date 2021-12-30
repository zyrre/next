package utils

import (
	"bufio"
	"fmt"
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

func CompleteTask(task string, mdMap map[string][]string) {
	file, err := os.Create("next.md")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	defer file.Close()
	startOfFile := true
	startOfCompleted := true
	for i, list := range mdMap {
		if !startOfFile {
			fmt.Fprintln(file, "")
		}
		startOfFile = false
		fmt.Fprintln(file, "# "+i)
		fmt.Fprintln(file, "")
		for _, line := range list {
			if line != task {
				if i == "Completed" {
					if startOfCompleted {
						fmt.Fprintln(file, "- ~~"+task+"~~")
						startOfCompleted = false
					}
					fmt.Fprintln(file, "- ~~"+line+"~~")
				} else {
					fmt.Fprintln(file, "- "+line)
				}
			}
		}
	}
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
