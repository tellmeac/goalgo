package app

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const ddl = `
create table if not exists chart(
    id integer primary key,
    timestamp integer not null,
    data text not null
);

create index if not exists chart_timestamp_idx on chart(timestamp); 
`

type Config struct{}

func New(_ *Config) *App {
	db, err := sqlx.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal(err)
	}

	_ = db.MustExec(ddl)

	return &App{db: db}
}

type App struct {
	db *sqlx.DB
}

// Chart contains aggregated info to be displayed as a chart.
type Chart struct {
	Data []Stamp `json:"stamps"`
}

type Stamp struct {
	Time        time.Time   `json:"x"`
	Candlestick CandleStick `json:"y"`
	// TODO: ... lines
}

type CandleStick struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	C float64 `json:"c"`
	D float64 `json:"d"`
}

func (a *App) GetLatest(from time.Time) (Chart, error) {
	panic("implement me")
}

func (a *App) GetInPeriod(from, to time.Time) (Chart, error) {
	panic("implement me")
}
