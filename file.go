package xoyutil

import (
	"errors"
	"io/ioutil"
	"os"
)

func fileRead(src string) string {
	content, err := ioutil.ReadFile(src)

	errorPanic(err)

	return string(content)
}

func fileAddLine(input string, filepath string) {
	_, err := os.Stat(filepath)
	if errors.Is(err, os.ErrNotExist) {
		fileCreate(filepath)
	}

	var file *os.File
	file, err = os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errorPanic(err)

	_, err = file.WriteString(input + "\n")
	errorPanic(err)
}

func fileCreate(filepath string) {
	_, err := os.Stat(filepath)
	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filepath)
		errorPanic(err)

		logger("fileCreate : file created -> " + filepath)
	}
}

func fileMkDir(folderpath string) {
	_, err := os.Stat(folderpath)
	if errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(folderpath, 0755)
		errorPanic(err)
		logger("fileMkDir : folder created -> " + folderpath)
	}
}
