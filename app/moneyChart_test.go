package app

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"graph/parser"
	"graph/parser/mocks"
	"testing"
)

func TestNewChart(t *testing.T) {
	p, ctrl := getParser(t)
	defer ctrl.Finish()
	mChart := NewChart(p, zerolog.Logger{})

	assert.Equal(t, 0, mChart.GetCollection().GetLength())
	mChart.Parse("test")
	fmt.Println(mChart.GetCollection())
	assert.Equal(t, 4, mChart.GetCollection().GetLength())
}

func TestMoneyChart_GetDatasets(t *testing.T) {
	p, ctrl := getParser(t)
	defer ctrl.Finish()
	mChart := NewChart(p, zerolog.Logger{})

	mChart.Parse("test")
	datasets := mChart.GetDatasets()

	assert.Equal(t, "income", datasets[0].Label)
	assert.Equal(t, []float64{100.86, 204, 90.93}, datasets[0].Values)
	//assert.Equal(t, 2, len(datasets))
	//fmt.Println(datasets)
	//mChart.
}

func TestMoneyChart_GetLabels(t *testing.T) {
	p, ctrl := getParser(t)
	defer ctrl.Finish()
	mChart := NewChart(p, zerolog.Logger{})
	assert.Equal(t, 0, len(mChart.GetLabels()))
	mChart.Parse("test")
	assert.Equal(t, 4, len(mChart.GetLabels()))
	assert.Equal(t, []string{
		"2023-04", "2023-05", "2023-06", "2023-07",
	}, mChart.GetLabels())
}

func getParser(t *testing.T) (parser.Parser, *gomock.Controller) {
	testData := map[int][]string{
		0: {"2023-04-15", "Кафе и рестораны", "TEST1", "", "счёт", "867,00", "RUB", "", "", "", "2023-04-15 22:23:49", "2023-04-17 05:12:42"},
		1: {"2023-05-16", "Кафе и рестораны1", "TEST2", "", "счёт", "100,00", "RUB", "", "", "", "2023-04-16 22:23:49", "2023-04-18 05:12:42"},
		2: {"2023-06-17", "Кафе и рестораны2", "TEST3", "", "счёт", "200,00", "RUB", "", "", "", "2023-04-17 22:23:49", "2023-04-19 05:12:42"},

		3: {"2023-04-10", "Зарплата", "test", "ЗП", "", "", "", "Текущий зарплатный счёт", "100,86", "RUB", "2023-04-10 12:28:39", "2023-04-12 17:28:28"},
		4: {"2023-05-10", "Зарплата", "test", "ЗП", "", "", "", "Текущий зарплатный счёт", "204", "RUB", "2023-04-10 12:28:39", "2023-04-12 17:28:28"},
		5: {"2023-06-10", "Зарплата", "test", "ЗП", "", "", "", "Текущий зарплатный счёт", "90,93", "RUB", "2023-04-10 12:28:39", "2023-04-12 17:28:28"},
	}

	ctrl := gomock.NewController(t)
	p := mocks.NewMockParser(ctrl)
	p.EXPECT().Parse("test").AnyTimes().Return(testData, nil)
	return p, ctrl
}
