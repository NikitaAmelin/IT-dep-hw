package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortLinesSlice(lines []string) []string {
	linesMap := make(map[string]int, len(lines))
	for _, i := range lines {
		linesMap[i]++
	}
	var uniqCount int
	for _, i := range linesMap {
		if i == 1 {
			uniqCount++
		}
	}
	uniqLines := make([]string, 0, uniqCount)
	for _, i := range lines {
		if linesMap[i] == 1 && i != "" {
			uniqLines = append(uniqLines, i)
		}
	}
	sort.Strings(uniqLines)
	return uniqLines
}

func makeLinesSlice(name string) ([]string, error) {
	fileBites, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	text := strings.ToUpper(string(fileBites))
	text = strings.Replace(text, "\r", "", -1)
	return strings.Split(text, "\n"), nil
}

func writeResult(name string) error {
	res, err := os.Create("result.txt")
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer res.Close()
	allLines, err := makeLinesSlice(name)
	if err != nil {
		return fmt.Errorf("ошибка считывания строк: %w", err)
	}
	uniqLines := sortLinesSlice(allLines)
	for i := range uniqLines {
		res.WriteString(fmt.Sprintf("%s - %d байт\n", uniqLines[i], len(uniqLines[i])))
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		panic(errors.New(`ошибка: не передано имя файла`))
	}
	name := os.Args[1]
	err := writeResult(name)
	if err != nil {
		panic(fmt.Errorf("ошибка записи в файл: %w", err))
	}
}
