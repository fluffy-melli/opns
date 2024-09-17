package Traffic

import "time"

type Count struct {
	Start   time.Time
	Event   int
	Slash   int
	Message int
}

func (c *Count) Reset() {
	c.Start = time.Now()
	c.Event = 0
	c.Slash = 0
	c.Message = 0
}

func (c *Count) Average_Traffic() float64 {
	tf := (c.Event + c.Slash + c.Message) / 3
	ti := time.Since(c.Start)
	return float64(tf) / float64(ti)
}
