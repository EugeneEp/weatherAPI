package alerts

const ServiceName = `AlertsServiceName`

type Alerts struct {
	msg  chan string
	done chan bool
}

func New() *Alerts {
	a := Alerts{
		msg: make(chan string),
	}
	go func() {
		for {
			select {
			case m := <-a.msg:
				println(m)
			case <-a.done:
				close(a.msg)
				break
			}
		}
	}()
	return &a
}

func (a *Alerts) Close() {
	a.done <- true
}

func (a *Alerts) Alert(msg string) {
	a.msg <- msg
}
