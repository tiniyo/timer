package timer

import (
	timer "github.com/singchia/go-timer"
)

type TiniyoTimer struct {
	m       map[string]timer.Tick
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
	tt.m = make(map[string]timer.Tick)
	tt.ch = make(chan interface{})
	tt.handler = cbHandler
	tt.t = timer.NewTimer()
	tt.t.Start()
}

func (this *TiniyoTimer) Run() {
	go func() {
		for {
			data, ok := <-this.ch
			if ok == false {
				break
			} else {
				this.handler(data.(timerdata).data)
				/* remove from map */
				delete(this.m, data.(timerdata).timer_id)
			}
		}
	}()
}

func (tt *TiniyoTimer) StartTimer(timerid string, data interface{}) error {
	tdata := timerdata{timer_id: timerid, data: data}
	t1, err := tt.t.Time(5, tdata, tt.ch, nil)
	if err != nil {
		return err
	}
	tt.m[timerid] = t1
	return nil
}

func (tt *TiniyoTimer) CancelTimer(msgid string) {
	tt.m[msgid].Cancel()
	/* remove from map */
	delete(tt.m, msgid)
}
