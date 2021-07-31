package midi

import (
	log "github.com/sirupsen/logrus"
	"gitlab.com/gomidi/midi"
	"gitlab.com/gomidi/rtmididrv"
)

type System struct {
	Driver *rtmididrv.Driver
	Ins    []midi.In
	Outs   []midi.Out
}

func NewSystem() *System {
	drv, err := rtmididrv.New()
	if err != nil {
		log.Panic(err)
	}

	ins, err := drv.Ins()
	if err != nil {
		log.Panic(err)
	}

	outs, err := drv.Outs()
	if err != nil {
		log.Panic(err)
	}

	return &System{
		Driver: drv,
		Ins:    ins,
		Outs:   outs,
	}
}

func (s *System) Close() {
	s.Driver.Close()
}

func MessageDispatch() {
	// Process the new messages defined here:
	// * https://github.com/erlsci/midilib/blob/release/0.4.x/src/midimsg.erl
	// that have been processed in pkg/erl/...?
}
