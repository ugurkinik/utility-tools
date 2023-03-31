package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func UpdateSelectedText(doAction func(input string) string, file string, selections []string) {
	// Read file line by line
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	lines := strings.Split(string(data), "\n")

	// Check each selection
	for i := len(selections) - 1; i >= 0; i-- {
		selectionParams := strings.Split(selections[i], ",")

		lineStart, _ := strconv.Atoi(selectionParams[0])
		columnStart, _ := strconv.Atoi(selectionParams[1])
		lineEnd, _ := strconv.Atoi(selectionParams[2])
		columnEnd, _ := strconv.Atoi(selectionParams[3])

		lineStart--
		lineEnd--

		// if the all selected characters are in the same lint
		if lineStart == lineEnd {
			selectedText := lines[lineStart][columnStart:columnEnd]
			result := doAction(selectedText)
			lines[lineStart] = fmt.Sprint(lines[lineStart][:columnStart], result, lines[lineStart][columnEnd:])
		} else { // or multiple lines are selected
			selectedText := fmt.Sprint(lines[lineStart][columnStart:], "\n"+strings.Join(lines[lineStart+1:lineEnd], "\n"), "\n", lines[lineEnd][:columnEnd])
			result := doAction(selectedText)
			lines[lineStart] = fmt.Sprint(lines[lineStart][:columnStart], result, lines[lineEnd][columnEnd:])
			lines = append(lines[:lineStart+1], lines[lineEnd+1:]...)
		}
	}

	// update the file
	updateFile, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
	}

	datawriter := bufio.NewWriter(updateFile)

	for lineNo, line := range lines {
		datawriter.WriteString(line)
		if lineNo < len(lines)-1 {
			datawriter.WriteByte(byte('\n'))
		}
	}

	datawriter.Flush()
	updateFile.Close()
}
