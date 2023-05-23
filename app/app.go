package app

import (
	"encoding/json"
	"net/http"
)

type App struct {
	ChartSrv *MoneyChart
}

func NewApp(ChartSrv *MoneyChart) App {
	return App{ChartSrv: ChartSrv}
}

func (app *App) GetChart(w http.ResponseWriter, r *http.Request) {

	if !app.ChartSrv.Parsed {
		app.ChartSrv.Parse("./zen_2023-05-23_dumpof_transactions_from_account9287301_alltime.csv")
	}

	w.Header().Set("Content-Type", "application/json")
	data := map[string]any{
		"labels":   app.ChartSrv.GetLabels(),
		"datasets": app.ChartSrv.GetDatasets(),
	}

	json.NewEncoder(w).Encode(data)
}
