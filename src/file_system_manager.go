package main

import (
	"os"
)

func writeToFileSystem(data []byte, _directory string, fileName string, fileExtension string, parentFolder string) (int, error) {
	if _, err := os.Stat(parentFolder + _directory + fileName + "." + fileExtension); err == nil {

		f, _ := os.OpenFile(parentFolder+_directory+fileName+"."+fileExtension, os.O_APPEND|os.O_WRONLY, 0644)
		_, writingError := f.Write(data)
		errorHandler(writingError)
		return 0, writingError

	}
	os.MkdirAll(parentFolder+_directory, os.ModePerm)
	fo, err := os.Create(parentFolder + _directory + fileName + "." + fileExtension)
	errorHandler(err)
	_, writingError := fo.Write(data)
	errorHandler(writingError)
	return 0, writingError
}
