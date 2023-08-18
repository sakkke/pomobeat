package main

import (
	"fmt"

	"github.com/sakkke/gopomobeat"
)

type viewModel struct {
	pomobeat  gopomobeat.Pomobeat
	eventType string
	duration  int
}

func newViewModel() *viewModel {
	pomobeat := gopomobeat.NewPomobeat()
	return &viewModel{
		pomobeat:  *pomobeat,
		eventType: getEventType(pomobeat),
		duration:  getDuration(pomobeat),
	}
}

func (v *viewModel) DecrementDuration() {
	v.duration--
}

func (v viewModel) GetDuration() int {
	return v.duration
}

func (v viewModel) GetEventType() string {
	return v.eventType
}

func (v viewModel) GetFormattedDuration() string {
	minutes := v.duration / 60
	seconds := v.duration % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func (v *viewModel) SetDuration(d int) {
	v.duration = d
}

func (v *viewModel) SetEventType(e string) {
	v.eventType = e
}

func (v *viewModel) SyncDuration() {
	v.SetDuration(getDuration(&v.pomobeat))
}

func (v *viewModel) SyncEventType() {
	v.SetEventType(getEventType(&v.pomobeat))
}

func getDuration(p *gopomobeat.Pomobeat) int {
	return int(p.GetDurationUntilNextEvent().Seconds())
}

func getEventType(p *gopomobeat.Pomobeat) string {
	texts := map[gopomobeat.EventType]string{
		gopomobeat.WorkTime:   "Work Time",
		gopomobeat.ShortBreak: "Short Break",
		gopomobeat.LongBreak:  "Long Break",
	}
	return texts[p.GetEventType()]
}
