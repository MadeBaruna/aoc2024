package lib

import (
	"bufio"
	"os"
)

func ReadFile(example bool) (string, error) {
	filepath := "input"
	if example {
		filepath += "_example"
	}

	file, err := os.Open(filepath + ".txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}
