package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManeger struct {
	InputFilePath  string
	OutputFilePath string
}

func New(input, output string) FileManeger {
	return FileManeger{
		InputFilePath:  input,
		OutputFilePath: output,
	}
}

func (f FileManeger) ReaLines() ([]string, error) {
	fp, err := os.Open(f.InputFilePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fp)
	var lines []string = make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fp.Close()
		return nil, err
	}
	fp.Close()
	return lines, nil
}

func (f FileManeger) WriteResult(data any) error {
	file, err := os.Create(f.OutputFilePath)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil
}
