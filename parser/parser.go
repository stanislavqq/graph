package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

type Parser interface {
	Parse(fileName string) map[int][]string
}

type CSVParser struct {
}

func NewParser(parserType string) (Parser, error) {
	if parserType == "csv" {
		return &CSVParser{}, nil
	}

	return nil, errors.New(fmt.Sprintf("Parser type \"%s\" is not available", parserType))
}

func (p *CSVParser) Parse(fileName string) map[int][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 12
	reader.Comment = '#'

	result := make(map[int][]string)

	i := 0
	for {
		record, e := reader.Read()
		if e != nil {

			if e.Error() == "EOF" {
				break
			}
			log.Default().Println(e)
			continue
		}

		result[i] = record
		i++
	}

	return result
}
