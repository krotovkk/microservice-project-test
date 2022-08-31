package counters

import (
	"strconv"
	"sync"
)

type Counter struct {
	cnt int
	M   *sync.RWMutex
}

func (c *Counter) Inc() {
	c.M.Lock()
	defer c.M.Unlock()
	c.cnt++
}

func (c *Counter) String() string {
	c.M.RLock()
	defer c.M.RUnlock()
	return strconv.FormatInt(int64(c.cnt), 10)
}
