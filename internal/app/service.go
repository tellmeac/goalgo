package app

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const ddl = `
create table if not exists chart(
    id integer primary key,
    timestamp integer not null,
    data text not null
);

create index if not exists chart_timestamp_idx on chart(timestamp); 
`

type Config struct {
	DatabaseConn string `yaml:"databaseConn"`
}

func New(c *Config) *Service {
	db, err := sqlx.Open("sqlite3", c.DatabaseConn)
	if err != nil {
		log.Fatal(err)
	}

	_ = db.MustExec(ddl)

	return &Service{db: db}
}

type Service struct {
	db *sqlx.DB
}

// Chart contains aggregated info to be displayed as a chart.
type Chart struct {
	Data []Stamp `json:"stamps"`
}

type Stamp struct {
	Time        int64       `json:"x"`
	Candlestick CandleStick `json:"y"`

	TopLine   float64  `json:"topLine"`
	DownLine  float64  `json:"downLine"`
	BlueLine  float64  `json:"blueLine"`
	NeedPoint *float64 `json:"needPoint"` // TODO: bool value
}

func (s *Stamp) Scan(val any) error {
	str, ok := val.(string)
	if !ok {
		return errors.New("invalid value")
	}

	if err := json.Unmarshal([]byte(str), s); err != nil {
		return err
	}

	return nil
}

type CandleStick struct {
	Open  float64 `json:"open"`
	High  float64 `json:"high"`
	Low   float64 `json:"low"`
	Close float64 `json:"close"`
}

type stampSQL struct {
	ID        int64 `db:"id"`
	Timestamp int64 `db:"timestamp"`
	Data      Stamp `db:"data"`
}

func (e *Service) GetLatest(ctx context.Context, from int64) (Chart, error) {
	args := map[string]any{"from": from}
	q := "select * from chart where timestamp > :from order by timestamp"

	rows, err := e.db.NamedQueryContext(ctx, q, args)
	if err != nil {
		return Chart{}, err
	}

	return rowsToChart(rows)
}

func (e *Service) GetInPeriod(ctx context.Context, from, to int64) (Chart, error) {
	args := map[string]any{
		"from": from,
		"to":   to,
	}
	q := "select * from chart where timestamp > :from and timestamp < :to order by timestamp"

	rows, err := e.db.NamedQueryContext(ctx, q, args)
	if err != nil {
		return Chart{}, err
	}

	return rowsToChart(rows)
}

func rowsToChart(rows *sqlx.Rows) (Chart, error) {
	data := make([]Stamp, 0)

	for rows.Next() {
		var dest stampSQL

		err := rows.StructScan(&dest)
		if err != nil {
			return Chart{}, err
		}

		data = append(data, dest.Data)
	}

	return Chart{Data: data}, nil
}