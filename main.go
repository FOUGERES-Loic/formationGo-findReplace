package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	src: source filename
	old/new: old word/new word
	occ: number of occurrences
	lines: slice with line numbers where 'old' was found
	err: function error
*/
func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		return occ, lines, err
	}
	defer dstFile.Close()

	old = old + " "
	new = new + " "
	scanner := bufio.NewScanner(srcFile)
	writter := bufio.NewWriter(dstFile)
	defer writter.Flush()

	line := 1
	for scanner.Scan() {
		found, res, occInLine := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += occInLine
			lines = append(lines, line)
		}
		fmt.Fprintln(writter, res)
		line++
	}

	return occ, lines, nil
}

/*
	line: line to process
	old/new: old word/new word
	found: true if at least one occurrence found
	res: new line with the word replaced (res == line if no change)
	occ: number of occurrences in said line
*/
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
	}
	return found, res, occ
}

func main() {
	old := "sit"
	new := "dot"
	occ, lines, err := FindReplaceFile("test.txt", "result.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing FindReplaceFile: %v\n", err)
	}
	fmt.Printf("number of occurrences of %v: %v\n", old, occ)
	fmt.Printf("number of lines: %d\n", len(lines))
	fmt.Print("Lines: [ ")
	len := len(lines)
	for i, l := range lines {
		fmt.Printf("%v", l)
		if i < len-1 {
			fmt.Print(" - ")
		}
	}
	fmt.Println(" ]")
}
