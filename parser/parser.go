package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
)

//go:generate mockgen -source=parser.go -destination mocks/parser.go -package mocks parser
type Parser interface {
	Parse(fileName string) (map[int][]string, error)
	SetSkipRows(num int) Parser
}

type CSVParser struct {
	SkipRowsCount int
}

func (p *CSVParser) SetSkipRows(num int) Parser {
	p.SkipRowsCount = num
	return p
}

func NewParser(parserType string) (Parser, error) {
	if parserType == "csv" {
		return &CSVParser{}, nil
	}

	return nil, errors.New(fmt.Sprintf("Parser type \"%s\" is not available", parserType))
}

func (p *CSVParser) Parse(fileName string) (map[int][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 12
	reader.Comment = '#'

	result := make(map[int][]string)

	slipCount := 0
	if p.SkipRowsCount > 0 {
		slipCount = p.SkipRowsCount
	}

	i := 0
	for {
		record, e := reader.Read()
		if slipCount > 0 {
			slipCount--
			continue
		}

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

	return result, nil
}
