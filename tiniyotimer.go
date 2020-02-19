package timer

import (
	"sync"

	timer "github.com/singchia/go-timer"
)

type TiniyoTimer struct {
	m       sync.Map
	ch      chan interface{}
	handler TiniyoTimerHandler
	t       timer.Timer
}

type timerdata struct {
	timer_id string
	data     interface{}
}

type TiniyoTimerHandler func(data interface{}) error

/* message id to timer tick */

func (tt *TiniyoTimer) InitializeTiniyoTimer(cbHandler TiniyoTimerHandler) {
	tt.ch = make(chan interface{})
	tt.handler = cbHandler
	tt.t = timer.NewTimer()
	tt.t.Start()
}

func (tt *TiniyoTimer) Run() {
	go func() {
		for {
			data, ok := <-tt.ch
			if ok == false {
				break
			} else {
				tt.handler(data.(timerdata).data)
				/* remove from map */
				tt.m.Delete(data.(timerdata).timer_id)
			}
		}
	}()
}

func (tt *TiniyoTimer) StartTimer(t uint64, timerid string, data interface{}) error {
	tdata := timerdata{timer_id: timerid, data: data}
	t1, err := tt.t.Time(t, tdata, tt.ch, nil)
	if err != nil {
		return err
	}
	tt.m.Store(timerid, t1)
	return nil
}

func (tt *TiniyoTimer) CancelTimer(timerid string) {
	timerobj, _ := tt.m.Load(timerid)
	timerobj.(timer.Tick).Cancel()
	/* remove from map */
	tt.m.Delete(timerid)
}
