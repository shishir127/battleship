package interfaces

import (
	"bufio"
	"errors"
	"os"
)

type FileInterface struct {
	InputFileName  string
	OutputFileName string
}

func (fileInterface *FileInterface) ReadInput() ([]string, error) {
	if fileInterface.InputFileName == "" {
		return nil, errors.New("Input file needs to be specified")
	}

	var input []string

	inFile, err := os.Open(fileInterface.InputFileName)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil

}
