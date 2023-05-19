package app

import (
	"fmt"
	"graph/money"
	"graph/parser"
	"log"
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

type Label string

type ChartList map[Label]ChartItem

type MoneyChart struct {
	Items *map[string]ChartList
}

func NewChart() *MoneyChart {
	mChart := &MoneyChart{}
	mChart.Items = &map[string]ChartList{}

	CSVParser, err := parser.NewParser("csv")
	if err != nil {
		log.Fatal(err)
	}
	items := CSVParser.Parse("file")
	list := ChartList{}

	for i := 1; i < len(items); i++ {
		row := items[i]
		moneyRow := money.NewFromRow(row)
		label := Label(moneyRow.Date.Format("2006-Jan"))

		cost := 0.0
		if moneyRow.IsIncome() {
			cost = moneyRow.Income
		} else if moneyRow.IsOutcome() {
			cost = moneyRow.Outcome
		}

		if val, ok := mChart.Items[moneyRow.GetType()][label]; ok {

			if moneyRow.GetType() == "income" && label == "2022-Dec" {
				fmt.Println(moneyRow.GetType()+string(label), cost, val.Value)
			}

			val.Value = val.Value + cost

		} else {
			list[label] = ChartItem{
				Name:  string(label),
				Value: cost,
				Type:  moneyRow.GetType(),
			}
			mChart.Items[moneyRow.GetType()] = list
		}
	}

	//fmt.Println(mChart.GetLabels())
	//fmt.Println(mChart.GetDatasets())

	return mChart
}

func (c *MoneyChart) GetLabels() []string {
	result := []string{}
	for _, item := range c.Items {
		for _, row := range item {
			result = append(result, row.Name)
		}
	}

	return result
}

func (c *MoneyChart) GetDatasets() []MoneyDataset {
	res := []MoneyDataset{}
	income := MoneyDataset{Label: "Income", Values: []float64{}}
	outcome := MoneyDataset{Label: "Outcome", Values: []float64{}}

	fmt.Println(c.Items)

	for _, item := range c.Items {
		for _, row := range item {
			fmt.Println(row.Type)
			if row.Type == TYPE_IN {
				income.Values = append(income.Values, row.Value)
			} else if row.Type == TYPE_OUT {
				outcome.Values = append(income.Values, row.Value)
			}
		}
	}

	fmt.Println(income.Values)
	fmt.Println(outcome.Values)

	res = append(res, income)
	res = append(res, outcome)

	return res
}
