package singleflight

import "sync"

// call is an in-flight or completed Do call.
type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	muLock sync.Mutex       // protects m
	m      map[string]*call //lazily initialized
}

// Do execute and returns the results of given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.muLock.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.muLock.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.muLock.Unlock()

	c.val, c.err = fn()
	c.wg.Done()

	g.muLock.Lock()
	delete(g.m, key)
	g.muLock.Unlock()

	return c.val, c.err
}
