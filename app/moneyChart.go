package app

import (
	"github.com/rs/zerolog"
	"graph/money"
	"graph/parser"
	"sort"
)

type ChartItem struct {
	Name  string
	Value float64
	Type  string
}

type MoneyDataset struct {
	Label  string
	Values []float64
}

const TYPE_OUT = "outcome"
const TYPE_IN = "income"
const DATE_FORMAT_DEFAULT = "2006-01"

type ChartList map[string]ChartItem

type MoneyChart struct {
	Labels     []string
	Datasets   ChartList
	Logger     zerolog.Logger
	Parser     parser.Parser
	collection *money.MoneyCollection
	Parsed     bool
}

func (chi *ChartItem) IncCost(value float64) {
	chi.Value += value
}

func NewChart(parser parser.Parser, logger zerolog.Logger) *MoneyChart {
	return &MoneyChart{
		Labels:     []string{},
		Datasets:   ChartList{},
		Logger:     logger,
		Parser:     parser,
		collection: money.NewCollection(),
		Parsed:     false,
	}
}

func (c *MoneyChart) GetCollection() *money.MoneyCollection {
	return c.collection
}

func (c *MoneyChart) Parse(file string) *MoneyChart {

	items, err := c.Parser.Parse(file)
	if err != nil {
		c.Logger.Error().Err(err).Msg("File parse error")
		panic(err)
	}

	c.collection = money.NewCollection()

	for _, item := range items {
		c.collection.Add(item)
	}
	c.Parsed = true
	return c
}

func (c *MoneyChart) GetLabels() []string {

	if len(c.Labels) > 0 {
		return c.Labels
	}

	if c.collection.GetLength() == 0 {
		c.Logger.Warn().Msg("Collection is empty. You can run parse func")
		return []string{}
	}

	keys := make(map[string]bool)

	c.collection.Each(func(key int, row money.MoneyRow) {
		label := row.Date.Format(DATE_FORMAT_DEFAULT)
		if _, value := keys[label]; !value {
			keys[label] = true
			c.Labels = append(c.Labels, label)
		}
	})

	sort.Strings(c.Labels)
	return c.Labels
}

func (c *MoneyChart) GetDatasets() []MoneyDataset {

	incomeDS := MoneyDataset{Label: TYPE_IN, Values: []float64{}}
	outcomeDS := MoneyDataset{Label: TYPE_OUT, Values: []float64{}}

	incomeMap := map[string]float64{}
	outcomeMap := map[string]float64{}

	c.collection.Each(func(key int, row money.MoneyRow) {
		if row.IsIncome() {
			incomeMap = addOrPlus(incomeMap, &row)
		}

		if row.IsOutcome() {
			outcomeMap = addOrPlus(outcomeMap, &row)
		}
	})

	for _, label := range c.GetLabels() {
		incomeDS.Values = append(incomeDS.Values, incomeMap[label])
		outcomeDS.Values = append(outcomeDS.Values, outcomeMap[label])
	}

	return []MoneyDataset{incomeDS, outcomeDS}
}

func addOrPlus(list map[string]float64, row *money.MoneyRow) map[string]float64 {
	label := row.Date.Format(DATE_FORMAT_DEFAULT)

	if _, exist := list[label]; exist {
		list[label] += row.GetCost()
	} else {
		list[label] = row.GetCost()
	}

	return list
}
