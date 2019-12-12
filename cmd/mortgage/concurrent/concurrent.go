package concurrent

import (
	"log"
	"sync"
)

var RequestsQueue = NewConcurrentSlice()

var StatusesQueue = NewConcurrentSlice()

// Concurrent map
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

// Concurrent slice
type ConcurrentSlice struct {
	sync.RWMutex
	items []interface{}
}

func NewConcurrentSlice() *ConcurrentSlice {
	return &ConcurrentSlice{items: make([]interface{}, 0)}

}

type ConcurrentSliceItem struct {
	Index int
	Value interface{}
}

func (cs *ConcurrentSlice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()

	cs.items = append(cs.items, item)
}

func (cs *ConcurrentSlice) Iter() <-chan ConcurrentSliceItem {
	c := make(chan ConcurrentSliceItem)

	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- ConcurrentSliceItem{index, value}
		}
		close(c)
	}
	go f()

	return c
}

// Queue methods implementation
func (cs *ConcurrentSlice) Enqueue(item interface{}) {
	cs.Append(item)
}

func (cs *ConcurrentSlice) Dequeue() {
	cs.Lock()
	defer cs.Unlock()

	cs.items = cs.items[1:]
}

func (cs *ConcurrentSlice) IsEmpty() bool {
	cs.Lock()
	defer cs.Unlock()

	return len(cs.items) == 0
}

func (cs *ConcurrentSlice) Peek() interface{} {
	cs.Lock()
	defer cs.Unlock()

	return cs.items[0]
}
