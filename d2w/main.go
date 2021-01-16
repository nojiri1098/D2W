package d2w

import (
    "bufio"
    "example.com/m/v2/util/file"
    "example.com/m/v2/util/logger"
    "fmt"
    "github.com/thoas/go-funk"
    "go.uber.org/zap"
    "os"
    "strings"
    "time"
)

const (
    inputFilename = "20060102.md"
    outputFilename = "%d/%d.%02d.md"
)

var (
    lines []file.Line
)

func D2W(year string) {
    filenames := funk.Filter(file.DirInfo(year), func(md string) bool {
        return !strings.HasPrefix(md, year + ".")
    }).([]string)

    for _, input := range filenames {
        output, err := buildOutputFilename(input)
        if err != nil {
            logger.Error("invalid input filename", zap.Error(err))
        }

        logger.Debug(fmt.Sprintf("start %s...", input))

        if lines, err = file.Read(year + "/" + input); err != nil {
            fmt.Println(os.Stderr, err)
            os.Exit(1)
        }

        if err := file.Write(output, year, lines, func(w *bufio.Writer, ls []file.Line) {
            for _, l := range lines {
                switch {
                case l.IsH1():
                    w.WriteString("#" + l.String()[:12]  + "\n")

                case l.IsTag():
                    // do nothing

                case l.IsWrongCodeBlock():
                    w.WriteString(l.CorrectCodeBlock() + "\n")

                case l.IsWrongQuote():
                    w.WriteString(l.CorrectQuote() + "\n")

                case l.IsWrongBundle():
                    w.WriteString(l.CorrectBundle() + "\n")
                    
                default:
                    w.WriteString(strings.TrimSpace(l.String()) + "\n")
                }
            }
        }); err != nil {
            fmt.Println(os.Stderr, err)
            os.Exit(1)
        }

        logger.Debug(fmt.Sprintf("end %s", input))
    }
}

func buildOutputFilename(input string) (string, error) {
    t, err := time.Parse(inputFilename, input)
    if err != nil {
        return "", err
    }

    y, d := t.ISOWeek()

    if err := file.MakeDirIfNotExists(fmt.Sprintf("%d", y)); err != nil {
        return "", err
    }

    return fmt.Sprintf(outputFilename, y, y, d), nil
}
