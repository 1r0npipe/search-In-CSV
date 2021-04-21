package initial

import (
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/1r0npipe/search-In-CSV/cmd/server/config"
	"github.com/1r0npipe/search-In-CSV/internal/model"
)

var (
	CSVSize      = 2048
	ErrLogFile   = errors.New("can't open/create log file")
	ErrErrorFile = errors.New("can't open/create error file")
	ErrCSVFile   = errors.New("can't open/create CSV file")
	ErrCSVRead   = errors.New("can't read from CSV file")
)

func FilesInit(c *config.Config) (*model.InitialModel, *os.File, *os.File, error) {
	mapka := make(model.InitialModel, CSVSize)
	logFile, err := LogFileInit(c)
	if err != nil {
		return nil, nil, nil, err
	}
	errorFile, err := ErrorFileInit(c)
	if err != nil {
		return nil, nil, nil, err
	}
	err = CSVFileRead(c, mapka)
	if err != nil {
		return nil, nil, nil, err
	}
	return &mapka, logFile, errorFile, nil
}

func LogFileInit(c *config.Config) (*os.File, error) {
	file, err := os.OpenFile(c.LogFile, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, ErrLogFile
	}
	defer file.Close()
	return file, nil
}

func ErrorFileInit(c *config.Config) (*os.File, error) {
	file, err := os.OpenFile(c.ErrorFile, os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, ErrLogFile
	}
	defer file.Close()
	return file, nil
}

func CSVFileRead(c *config.Config, m model.InitialModel) error {
	csvFile, err := os.OpenFile(c.DefaultFile, os.O_RDONLY, 0666)
	if err != nil {
		return ErrCSVFile
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	//reader.Comma = ','
	header, err := reader.Read()
	if err != nil {
		return ErrCSVRead
	}
	m[0] = header
	for i := 1; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		m[i] = record
	}
	return nil
}

// func (f *os.File) Read(b []byte) (int, error) {

// }
