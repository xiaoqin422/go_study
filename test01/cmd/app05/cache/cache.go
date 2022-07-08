package cache

import "sync"

type cache struct {
	f    Func
	memo map[string]*entry
	mu   sync.Mutex // 重复抑制
}
type Func func(name string) (interface{}, error)
type entry struct {
	res   result
	ready chan struct{}
}
type result struct {
	value interface{}
	err   error
}

func New(f Func) *cache {
	return &cache{f: f, memo: make(map[string]*entry)}
}
func (c *cache) Get(name string) (interface{}, error) {
	c.mu.Lock()
	e := c.memo[name]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		c.memo[name] = e
		c.mu.Unlock()
		e.res.value, e.res.err = c.f(name)
		close(e.ready)
	} else {
		c.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
