package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filename = "next.md"

func FileToMap(filename string) (map[string][]string, error) {
	lines, _, _, err := parseFile(filename)
	if err != nil {
		return nil, err
	}
	return mdDataToMap(lines), nil
}

func CompleteTask() (string, error) {
	lines, doingStart, completedStart, err := parseFile(filename)
	if err != nil {
		return "", err
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return "", err
	}
	defer file.Close()
	for i, line := range lines {
		if doingStart != 0 && i == doingStart+2 {
			//line of completed task, so not printing it under the "Doing" section
		} else if completedStart != 0 && i == completedStart+2 {
			fmt.Fprintln(file, "- ~~"+strings.Trim(lines[doingStart+2], "- ")+"~~")
			fmt.Fprintln(file, line)
		} else {
			fmt.Fprintln(file, line)
		}
	}
	return "", nil
}

func UndoTask() (string, error) {
	lines, doingStart, completedStart, err := parseFile(filename)
	if err != nil {
		return "", err
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return "", err
	}
	defer file.Close()
	for i, line := range lines {
		if doingStart != 0 && i == doingStart+2 {
			fmt.Fprintln(file, "- "+strings.Trim(lines[completedStart+2], "- ~"))
			fmt.Fprintln(file, line)
		} else if completedStart != 0 && i == completedStart+2 {
			//line of task to be undone, so not printing it under the "Completed" section
		} else {
			fmt.Fprintln(file, line)
		}
	}
	return "", nil
}

func mdDataToMap(mdData []string) map[string][]string {
	mdMap := make(map[string][]string)
	var currentIndex string
	for _, line := range mdData {
		line := strings.TrimSpace(line)
		if line != "" {
			splitLine := strings.SplitN(line, " ", 2)
			if splitLine[0] == "#" {
				mdMap[splitLine[1]] = []string{}
				currentIndex = splitLine[1]
			} else {
				mdMap[currentIndex] = append(mdMap[currentIndex], strings.Trim(splitLine[1], "~"))
			}
		}
	}
	return mdMap
}

//parse a file and return each non-whitespace line
func parseFile(filename string) ([]string, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	doingStart := 0
	completedStart := 0
	for scanner.Scan() {
		if scanner.Text() == "# Doing" {
			doingStart = len(lines)
		} else if scanner.Text() == "# Completed" {
			completedStart = len(lines)
		}
		lines = append(lines, scanner.Text())
	}
	return lines, doingStart, completedStart, scanner.Err()
}
