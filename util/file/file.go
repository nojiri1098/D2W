package file

import (
	"bufio"
	"github.com/thoas/go-funk"
	"os"
	"strings"
)

const (
	prefixH1  = "# "
	prefixTag = "#"
	prefixWrongCodeBlock = "`  `  `"
	prefixCodeBlock = "```"
	prefixWrongQuote = " >"
	prefixQuote = ">"
	prefixWrongBundle = "・"
	prefixBundle = "- "
)

type Line string

func (l Line) IsH1() bool {
	return strings.HasPrefix(l.String(), prefixH1)
}

func (l Line) IsTag() bool {
	return !l.IsH1() && strings.HasPrefix(l.String(), prefixTag)
}

func (l Line) IsWrongCodeBlock() bool {
	return strings.HasPrefix(l.String(), prefixWrongCodeBlock)
}

func (l Line) CorrectCodeBlock() string {
	if l.IsWrongCodeBlock() {
		return prefixCodeBlock + strings.TrimPrefix(l.String(), prefixWrongCodeBlock)
	}

	return l.String()
}

func (l Line) IsWrongQuote() bool {
	return strings.HasPrefix(l.String(), prefixWrongQuote)
}

func (l Line) CorrectQuote() string {
	if l.IsWrongQuote() {
		return prefixQuote + strings.Trim(l.String(), prefixWrongQuote)
	}

	return l.String()
}

func (l Line) IsWrongBundle() bool {
	return strings.HasPrefix(l.String(), prefixWrongBundle)
}

func (l Line) CorrectBundle() string {
	if l.IsWrongBundle() {
		return prefixBundle + strings.Trim(l.String(), prefixWrongBundle)
	}

	return l.String()
}

func (l Line) String() string { return string(l) }

func List(dir string, handler func(string) bool) []string {
	return funk.Filter(DirInfo(dir), handler).([]string)
}

func Rename(old, new string) error {
	return os.Rename(old, new)
}

func Read(filename string) ([]Line, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, Line(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

type WriteHandler func(*bufio.Writer, []Line)

func Write(filename string, year string, lines []Line, handler WriteHandler) error {
	_, err := os.Stat(filename)
	isFileExists := err == nil

	fp, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()

	writer := bufio.NewWriter(fp)

	if !isFileExists {
		h1Str := strings.TrimSuffix(strings.TrimPrefix(filename, year + "/"), ".md")
		writer.WriteString(prefixH1 + h1Str + "\n\n")
	}

	handler(writer, lines)

	writer.WriteString("\n")
	writer.Flush()

	lines = []Line{}

	return nil
}
