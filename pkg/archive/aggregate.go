package archive

import (
	"fmt"
	"log"
	"time"

	"github.com/go-xorm/xorm"
)

// get X events from DB.
// Parse add to memory
// loop over and aggregate per day.

type AggregatedStats struct {
	IssueCount int64
}

type Aggregator struct {
	engine *xorm.Engine
}

func NewAggregator(engine *xorm.Engine) *Aggregator {
	return &Aggregator{engine: engine}
}

func (a *Aggregator) Aggregate() error {
	events := []*GithubEvent{}

	err := a.engine.Limit(1000, 0).Find(&events)
	if err != nil {
		log.Fatalf("failed to query. error: %v", err)
	}

	for _, e := range events {
		issueStat, err := e.Payload.Get("action").String()
		if err == nil && issueStat != "" {
			fmt.Printf("date: %v issue action: %s\n", e.CreatedAt, issueStat)
		}
	}

	return nil
}

func GetArchDateFrom(t time.Time) *ArchiveFile {
	return NewArchiveFile(t.Year(), int(t.Month()), t.Day(), t.Hour())
}
