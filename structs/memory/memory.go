package memory

import (
	"github.com/ungame/go-generics-sandbox/structs/models"
	"sync"
)

type Memory interface {
	*models.User | *models.Comment
}

type memoryImpl[T Memory] struct {
	data map[string]T
	lk   *sync.Mutex
}

func New[T Memory]() *memoryImpl[T] {
	return &memoryImpl[T]{
		data: make(map[string]T),
		lk:   &sync.Mutex{},
	}
}

func (m *memoryImpl[T]) onLock(rw func()) {
	m.lk.Lock()
	rw()
	m.lk.Unlock()
}

func (m *memoryImpl[T]) Set(key string, val T) {
	m.onLock(func() {
		m.data[key] = val
	})
}

func (m *memoryImpl[T]) Get(key string) (val T) {
	m.onLock(func() {
		val = m.data[key]
	})
	return
}
