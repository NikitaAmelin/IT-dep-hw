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

func writeFile(name string, uniqLines []string) error {
	res, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer res.Close()
	for i := range uniqLines {
		res.WriteString(fmt.Sprintf("%s - %d байт\n", uniqLines[i], len(uniqLines[i])))
	}
	return nil
}

func main() {
	if len(os.Args) < 3 {
		panic(errors.New(`ошибка: переданы не все файлы (исходный файл, файл вывода)`))
	}
	input_file, output_file := os.Args[1], os.Args[2]
	lines, err := makeLinesSlice(input_file)
	if err != nil {
		panic(fmt.Errorf(`ошибка при попытке считать строки: %w`, err))
	}
	uniqLines := sortLinesSlice(lines)
	err = writeFile(output_file, uniqLines)
	if err != nil {
		panic(fmt.Errorf("ошибка записи в файл: %w", err))
	}
}
