package money

import (
	"time"
)

type MoneyRow struct {
	Date                      time.Time
	CategoryName              string
	Payee                     string
	Comment                   string
	OutcomeAccountName        string
	Outcome                   float64
	OutcomeCurrencyShortTitle string
	IncomeAccountName         string
	Income                    float64
	IncomeCurrencyShortTitle  string
	CreatedDate               time.Time
	ChangedDate               time.Time
}

func New() MoneyRow {
	return MoneyRow{}
}

func NewFromRow(row []string) MoneyRow {
	moneyRow := New()

	for index, val := range row {
		switch index {
		case 0:
			moneyRow.SetDate(val)
		case 1:
			moneyRow.SetCategoryName(val)
		case 2:
			moneyRow.SetPayee(val)
		case 3:
			moneyRow.SetComment(val)
		case 4:
			moneyRow.SetOutcomeAccountName(val)
		case 5:
			moneyRow.SetOutcome(val)
		case 6:
			moneyRow.SetOutcomeCurrencyShortTitle(val)
		case 7:
			moneyRow.SetIncomeAccountName(val)
		case 8:
			moneyRow.SetIncome(val)
		case 9:
			moneyRow.SetIncomeCurrencyShortTitle(val)
		case 10:
			moneyRow.SetCreatedDate(val)
		case 11:
			moneyRow.SetChangedDate(val)
		default:
			panic("default call")
		}
	}

	return moneyRow
}
func (m *MoneyRow) SetOutcomeAccountName(value string) *MoneyRow {
	m.OutcomeAccountName = value
	return m
}

func (m *MoneyRow) SetOutcome(value string) *MoneyRow {
	m.Outcome = parseCostValue(value)
	return m
}

func (m *MoneyRow) SetOutcomeCurrencyShortTitle(value string) *MoneyRow {
	m.OutcomeCurrencyShortTitle = value
	return m
}

func (m *MoneyRow) SetIncomeAccountName(value string) *MoneyRow {
	m.IncomeAccountName = value
	return m
}

func (m *MoneyRow) SetIncome(value string) *MoneyRow {
	m.Income = parseCostValue(value)
	return m
}

func (m *MoneyRow) SetIncomeCurrencyShortTitle(value string) *MoneyRow {
	m.IncomeCurrencyShortTitle = value
	return m
}

func (m *MoneyRow) SetCreatedDate(value string) *MoneyRow {
	date, err := time.Parse("2006-01-02 15:04:05", value)
	if err != nil {
		panic(err)
	}

	m.CreatedDate = date
	return m
}

func (m *MoneyRow) SetChangedDate(value string) *MoneyRow {
	date, err := time.Parse("2006-01-02 15:04:05", value)
	if err != nil {
		panic(err)
	}

	m.ChangedDate = date
	return m
}

func (m *MoneyRow) SetComment(value string) *MoneyRow {
	m.Comment = value
	return m
}

func (m *MoneyRow) SetPayee(value string) *MoneyRow {
	m.Payee = value
	return m
}

func (m *MoneyRow) SetCategoryName(value string) *MoneyRow {
	m.CategoryName = value
	return m
}

func (m *MoneyRow) SetDate(value string) *MoneyRow {
	date, err := time.Parse("2006-01-02", value)
	if err != nil {
		panic(err)
	}

	m.Date = date
	return m
}

func (m *MoneyRow) IsIncome() bool {
	return m.Outcome == 0 && m.Income > 0
}

func (m *MoneyRow) IsOutcome() bool {
	return m.Outcome > 0 && m.Income == 0
}

func (m *MoneyRow) GetType() string {

	if m.IsIncome() {
		return "income"
	}

	if m.IsOutcome() {
		return "outcome"
	}

	return "unknown"
}

func parseCostValue(cost string) float64 {
	res, err := NewParser().ParseFloat(cost)
	if err != nil {
		panic(err)
	}

	return res
}
