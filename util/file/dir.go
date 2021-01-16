package file

import (
	"io/ioutil"
	"os"
	"strings"
)

func DirInfo(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, DirInfo(file.Name())...)
			continue
		}

		// 隠しファイルは無視
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}

		paths = append(paths, file.Name())
	}

	return paths
}

func MakeDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	return nil
}
