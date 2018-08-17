package main

import (
	"os"
)

func writeToFileSystem(data []byte, _directory string, fileName string, fileExtension string, parentFolder string) (int, error) {

	os.MkdirAll(parentFolder+_directory, os.ModePerm)
	fo, err := os.Create(parentFolder + _directory + fileName + "." + fileExtension)
	errorHandler(err)
	_, writingError := fo.Write(data)
	errorHandler(writingError)
	return 0, writingError
}
