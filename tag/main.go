package tag

import (
	"bufio"
	"example.com/m/v2/util/file"
	"example.com/m/v2/util/logger"
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"strings"
)

var (
	lines []file.Line
)

func Tag(year string) {
	filenames := funk.Filter(file.DirInfo(year), func(md string) bool {
		return strings.HasPrefix(md, year + ".")
	}).([]string)

	for _, fn := range filenames {
		logger.Debug(fmt.Sprintf("start %s...", fn))

		var err error
		if lines, err = file.Read(year + "/" + fn); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}

		output := fmt.Sprintf("%s/%s", year, fn)

		if err = file.Write(output, year, lines, func(w *bufio.Writer, _ []file.Line) {
			w.WriteString(fmt.Sprintf("#%s#", year))
		}); err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}

		logger.Debug(fmt.Sprintf("end %s", fn))
	}
}
