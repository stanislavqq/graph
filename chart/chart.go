package chart

import (
	"errors"
	"fmt"
)

type Chart struct {
	Type     string
	Labels   []string
	Datasets []Dataset
}

type Color struct {
	Value string
}

type Dataset struct {
	Label           string
	Value           int
	Border          int
	BorderColor     Color
	BackgroundColor Color
}

func NewChart(chartType string) *Chart {
	if chartType == "line" {
		return &Chart{Type: chartType}
	}

	panic(errors.New(fmt.Sprintf("Chart type \"%s\" is not available", chartType)))
}
