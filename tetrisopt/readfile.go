package tetrisopt

import (
	"bufio"
	"os"
)

func readFile(fileName string) (*[]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = append(lines, "")
	return &lines, nil
}
