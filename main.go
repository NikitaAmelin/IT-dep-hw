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
		return []string{}, fmt.Errorf("ошибка чтения файла: %v", err)
	}
	text := strings.ToUpper(string(fileBites))
	text = strings.Replace(text, "\r", "", -1)
	return strings.Split(text, "\n"), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(errors.New(`передайте имя файла`))
		return
	}
	name := os.Args[1]
	res, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("ошибка создания файла: %v", err))
		return
	}
	defer res.Close()
	allLines, err := makeLinesSlice(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	uniqLines := sortLinesSlice(allLines)
	for i := range uniqLines {
		res.WriteString(fmt.Sprintf("%s - %d байт\n", uniqLines[i], len(uniqLines[i])))
	}
}
