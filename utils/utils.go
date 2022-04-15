package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "next.md"

func InitFile() error {
	wd, _ := os.Getwd()
	_, err := os.Stat(wd + "/" + filename)
	if err == nil {
		return fmt.Errorf("file already exists")
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return err
	}
	defer file.Close()
	fmt.Fprintln(file, "# To do")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "# Doing")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "# Completed")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "")
	return nil
}
func FileToMap(filename string) (map[string][]string, error) {
	lines, _, _, _, _, err := parseFile(filename, "")
	if err != nil {
		return nil, err
	}
	return mdDataToMap(lines), nil
}

func CompleteTask() (string, error) {
	lines, doingStart, completedStart, _, _, err := parseFile(filename, "")
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

func UndoTask(taskName string) (string, error) {
	lines, doingStart, completedStart, taskLine, taskName, err := parseFile(filename, taskName)
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
		if i == taskLine && taskLine != 0 {
			//line of task to undo, so not printing it
		} else if doingStart != 0 && i == doingStart+2 && taskLine != 0 && taskLine > completedStart {
			fmt.Fprintln(file, "- "+taskName)
			fmt.Fprintln(file, line)
		} else if i == 3 && taskLine != 0 && taskLine < completedStart {
			fmt.Fprintln(file, "- "+taskName)
			fmt.Fprintln(file, line)
		} else {
			fmt.Fprintln(file, line)
		}
	}
	return "", nil
}

func AddTask(task string) error {
	lines, _, _, _, _, err := parseFile(filename, "")
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return err
	}
	defer file.Close()
	for i, line := range lines {
		if i == 2 {
			fmt.Fprintln(file, "- "+task)
		}
		fmt.Fprintln(file, line)
	}
	return nil
}

func RemoveTask(task string) error {
	lines, _, _, _, _, err := parseFile(filename, "")
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return err
	}
	defer file.Close()
	for _, line := range lines {
		if strings.Trim(line, "- ~") != task {
			fmt.Fprintln(file, line)
		}
	}
	return nil
}

func StartTask(task string) error {
	var taskInt int
	var taskName string
	if len(task) == 2 {
		taskInt, _ = strconv.Atoi(string(task[1]))
		taskInt = taskInt + 2
		task = strconv.Itoa(taskInt)
	}
	lines, doingStart, _, _, _, err := parseFile(filename, "")
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return err
	}
	defer file.Close()
	for i, line := range lines {
		if strconv.Itoa(i) == task {
			taskName = strings.Trim(line, "- ")
		}
		if doingStart != 0 && i == doingStart+2 {
			if taskName != "" {
				fmt.Fprintln(file, "- "+taskName)
			} else {
				fmt.Fprintln(file, "- "+task)
			}
		}
		if i > doingStart || (strings.Trim(line, "- ~") != task && strconv.Itoa(i) != task) {
			fmt.Fprintln(file, line)
		}
	}
	return nil
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
func parseFile(filename string, task string) ([]string, int, int, int, string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, 0, task, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	doingStart := 0
	completedStart := 0
	taskLine := 0
	taskLineOffset := 999
	taskArea := "nope"
	if len(task) == 2 {
		taskLineOffset, _ = strconv.Atoi(string(task[1]))
		taskLineOffset = taskLineOffset + 2
		taskArea = string(task[0])
	}
	currentTaskArea := ""
	i := 0
	for scanner.Scan() {
		switch scanner.Text() {
		case "# To do":
			currentTaskArea = "t"
			i = 0
		case "# Doing":
			currentTaskArea = "d"
			i = 0
		case "# Completed":
			currentTaskArea = "c"
			i = 0
		}
		if taskLineOffset != 999 && i == taskLineOffset && currentTaskArea == taskArea {
			task = strings.Trim(scanner.Text(), "- ~")
			taskLine = len(lines)
		}
		if scanner.Text() == "# Doing" {
			currentTaskArea = "d"
			i = 0
			doingStart = len(lines)
		} else if scanner.Text() == "# Completed" {
			completedStart = len(lines)
		} else if task != "" && strings.Trim(scanner.Text(), "- ~") == task {
			taskLine = len(lines)
		}
		lines = append(lines, scanner.Text())
		i++
	}
	return lines, doingStart, completedStart, taskLine, task, scanner.Err()
}
