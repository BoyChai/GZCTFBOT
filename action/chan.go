package action

type chans []chan string

var Chans chans

func init() {
	Chans = make(chans, 0)
}

func newChan() chan string {
	c := make(chan string)
	Chans = append(Chans, c)
	return c
}
