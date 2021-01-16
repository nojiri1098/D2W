package rename

import (
	"example.com/m/v2/util/file"
	"example.com/m/v2/util/logger"
	"fmt"
	"go.uber.org/zap"
	"strings"
)

func Rename(year string) {
	filenames := file.List(year, func(md string) bool {
		return !strings.HasPrefix(md, year + ".")
	})

	for _, old := range filenames {
		logger.Debug(fmt.Sprintf("start %s...", old))

		if err := file.Rename(
			fmt.Sprintf("%s/%s", year, old),
			fmt.Sprintf("%s/%s.md", year, old[:8]),
		); err != nil {
			logger.Error("rename error", zap.Error(err))
		}

		logger.Debug(fmt.Sprintf("end %s", old))
	}
}
