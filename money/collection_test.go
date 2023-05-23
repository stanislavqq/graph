package money

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCollection(t *testing.T) {
	collection := NewCollection()

	testData := map[int][]string{
		0: {"2023-04-15", "Кафе и рестораны", "TEST1", "", "счёт", "867,00", "RUB", "", "", "", "2023-04-15 22:23:49", "2023-04-17 05:12:42"},
		1: {"2023-04-16", "Кафе и рестораны1", "TEST2", "", "счёт", "100,00", "RUB", "", "", "", "2023-04-16 22:23:49", "2023-04-18 05:12:42"},
		2: {"2023-04-17", "Кафе и рестораны2", "TEST3", "", "счёт", "200,00", "RUB", "", "", "", "2023-04-17 22:23:49", "2023-04-19 05:12:42"},
		3: {"2023-04-18", "Кафе и рестораны3", "TEST4", "", "счёт", "342,21", "RUB", "", "", "", "2023-04-18 22:23:49", "2023-04-20 05:12:42"},
	}

	for i := 0; i < len(testData); i++ {
		list := testData[i]
		collection.Add(list)
	}
	date, _ := time.Parse("2006-01-02", "2023-04-15")
	created, _ := time.Parse("2006-01-02 15:04:05", "2023-04-15 22:23:49")
	updated, _ := time.Parse("2006-01-02 15:04:05", "2023-04-17 05:12:42")

	expected := MoneyRow{
		ID:                        0,
		Date:                      date,
		CategoryName:              "Кафе и рестораны",
		Payee:                     "TEST1",
		Comment:                   "",
		OutcomeAccountName:        "счёт",
		Outcome:                   867.00,
		OutcomeCurrencyShortTitle: "RUB",
		CreatedDate:               created,
		ChangedDate:               updated,
		Cost:                      867.00,
	}
	item, err := collection.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, expected, item)

	assert.Equal(t, 4, collection.GetLength())
}

func TestMoneyCollection_Add(t *testing.T) {
	collection := NewCollection()

	assert.Equal(t, 0, collection.GetLength())
	collection.Add([]string{"2023-04-15", "Кафе и рестораны", "TEST1", "", "счёт", "867,00", "RUB", "", "", "", "2023-04-15 22:23:49", "2023-04-17 05:12:42"})
	assert.Equal(t, 1, collection.GetLength())
	item, err := collection.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "2023-04-15", item.Date.Format("2006-01-02"))
	assert.Equal(t, "Кафе и рестораны", item.CategoryName)
	assert.Equal(t, "TEST1", item.Payee)
	assert.Equal(t, 867.00, item.GetCost())
}
