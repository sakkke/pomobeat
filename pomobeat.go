package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sakkke/gopomobeat"
)

func main() {
	vm := newViewModel()

	a := app.New()
	w := a.NewWindow("Pomobeat")

	eventTypeLabel := widget.NewLabel(vm.GetEventType())
	durationLabel := widget.NewLabel(vm.GetFormattedDuration())

	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			vm.DecrementDuration()
			durationLabel.SetText(vm.GetFormattedDuration())
		}
	}()

	listener := gopomobeat.EventListener(func(p gopomobeat.Pomobeat) {
		vm.pomobeat.Sync()

		vm.SyncEventType()
		eventTypeLabel.SetText(vm.GetEventType())

		vm.SyncDuration()
		durationLabel.SetText(vm.GetFormattedDuration())
	})

	vm.pomobeat.AddEventListener(gopomobeat.WorkTime, listener)
	vm.pomobeat.AddEventListener(gopomobeat.ShortBreak, listener)
	vm.pomobeat.AddEventListener(gopomobeat.LongBreak, listener)

	go vm.pomobeat.Listen()

	vbox := container.NewVBox(eventTypeLabel, durationLabel)
	w.SetContent(vbox)
	w.ShowAndRun()
}
