package common

import (
	"log"
	"sync"
)

type ConcurrentMap struct {
	mx  sync.Mutex
	Map map[string]interface{}
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		Map: make(map[string]interface{}),
	}
}

func (cm *ConcurrentMap) Get(key string) (interface{}, bool) {
	cm.mx.Lock()
	defer cm.mx.Unlock()
	val, ok := cm.Map[key]
	return val, ok
}

func (cm *ConcurrentMap) Add(key string, val interface{}) {
	cm.mx.Lock()
	defer cm.mx.Unlock()
	cm.Map[key] = val
}

func (cm *ConcurrentMap) Delete(key string) {
	cm.mx.Lock()
	defer cm.mx.Unlock()
	delete(cm.Map, key)
}

func (cm *ConcurrentMap) Iter() chan interface{} {

	c := make(chan interface{})

	go func() {
		cm.mx.Lock()
		defer cm.mx.Unlock()
		for _, v := range cm.Map {
			c <- v
		}
		close(c)
	}()

	return c
}

func (cm *ConcurrentMap) Print() {
	cm.mx.Lock()
	defer cm.mx.Unlock()
	log.Println(cm.Map)
}

type ConcurrentCounter struct {
	mx      sync.Mutex
	Counter int
}

func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{
		Counter: 0,
	}
}

func (cc *ConcurrentCounter) Inc() {
	cc.mx.Lock()
	defer cc.mx.Unlock()
	cc.Counter += 1
}

func (cc *ConcurrentCounter) Dec() {
	cc.mx.Lock()
	defer cc.mx.Unlock()
	cc.Counter -= 1
}

func (cc *ConcurrentCounter) Val() int {
	cc.mx.Lock()
	defer cc.mx.Unlock()
	return cc.Counter
}

func (cc *ConcurrentCounter) SetVal(val int) {
	cc.mx.Lock()
	defer cc.mx.Unlock()
	cc.Counter = val
}
