package money

import (
	"strconv"
	"strings"
)

type RubParser struct {
}

func NewParser() *RubParser {
	return &RubParser{}
}

func (p *RubParser) ParseFloat(value string) (float64, error) {
	var res float64
	var err error
	res = 0

	if len(value) == 0 {
		return res, nil
	}

	splited := strings.Split(value, ",")

	if len(splited) < 2 {
		parsed, err := strconv.ParseFloat(value, 64)
		return parsed, err
	}

	first, err := strconv.ParseFloat(splited[0], 64)
	second, err := strconv.ParseFloat("0."+splited[1], 64)
	res = first + second
	return res, err
}
