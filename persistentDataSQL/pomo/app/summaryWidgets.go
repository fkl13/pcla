package app

import (
	"context"
	"math"
	"time"

	"github.com/fkl13/pcla/persistentDataSQL/pomo/pomodoro"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/widgets/barchart"
	"github.com/mum4k/termdash/widgets/linechart"
)

type summary struct {
	bcDay        *barchart.BarChart
	lcWeekly     *linechart.LineChart
	updateDaily  chan bool
	updateWeekly chan bool
}

func (s *summary) update(redrawCh chan<- bool) {
	s.updateDaily <- true
	s.updateWeekly <- true
	redrawCh <- true
}

func newSummary(ctx context.Context, config *pomodoro.IntervalConfig, redraw chan<- bool, errorCh chan<- error) (*summary, error) {
	s := &summary{}
	var err error

	s.updateDaily = make(chan bool)
	s.updateWeekly = make(chan bool)

	s.bcDay, err = newBarChart(ctx, config, s.updateDaily, errorCh)
	if err != nil {
		return nil, err
	}

	s.lcWeekly, err = newLineChart(ctx, config, s.updateWeekly, errorCh)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func newBarChart(ctx context.Context, config *pomodoro.IntervalConfig,
	update <-chan bool, errorCh chan<- error,
) (*barchart.BarChart, error) {
	bc, err := barchart.New(
		barchart.ShowValues(),
		barchart.BarColors([]cell.Color{
			cell.ColorBlue,
			cell.ColorYellow,
		}),
		barchart.ValueColors([]cell.Color{
			cell.ColorBlack,
			cell.ColorBlack,
		}),
		barchart.Labels([]string{
			"Pomodoro",
			"Break",
		}),
	)
	if err != nil {
		return nil, err
	}

	updateWidget := func() error {
		ds, err := pomodoro.DailySummary(time.Now(), config)
		if err != nil {
			return err
		}

		return bc.Values(
			[]int{
				int(ds[0].Minutes()),
				int(ds[1].Minutes()),
			},
			int(math.Max(ds[0].Minutes(), ds[1].Minutes())*1.1)+1,
		)
	}

	// Update goroutine for BarChart
	go func() {
		for {
			select {
			case <-update:
				errorCh <- updateWidget()
			case <-ctx.Done():
				return
			}
		}
	}()

	// Force Update BarChart at start
	if err := updateWidget(); err != nil {
		return nil, err
	}

	return bc, nil
}

func newLineChart(ctx context.Context, config *pomodoro.IntervalConfig,
	update chan bool, errorCh chan<- error,
) (*linechart.LineChart, error) {
	panic("unimplemented")
}
