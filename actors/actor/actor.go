package actor

import (
	"fmt"
	"io"
	"sync"
)

// Event type of a event
type Event int

// Arg an argument of an event
type Arg interface{}

type eventEntity struct {
	event Event
	args  []Arg
}

// EventHandler handler of a type of event
type EventHandler func(args ...Arg)

// Actor interface of an actor
type Actor interface {
	io.Closer

	// Register register an event and its handler
	Register(event Event, handler EventHandler) error

	// Cast casts an event and its arguments
	Cast(event Event, args ...Arg) error
}

type actor struct {
	sync.RWMutex

	closeOnce sync.Once
	eventMap  map[Event]EventHandler

	eventCh chan eventEntity
	quit    chan struct{}
	done    chan struct{}
}

// New creates an actor
func New() Actor {
	a := &actor{
		eventMap: map[Event]EventHandler{},
		eventCh:  make(chan eventEntity),
		quit:     make(chan struct{}),
		done:     make(chan struct{}),
	}

	go a.process()

	return a
}

// Register register an event and its handler
func (a *actor) Register(event Event, handler EventHandler) error {
	a.Lock()
	defer a.Unlock()

	if handler == nil {
		return fmt.Errorf("handler is required")
	}

	if _, ok := a.eventMap[event]; ok {
		return fmt.Errorf("eventEntity %v has already registered", event)
	}

	a.eventMap[event] = handler
	return nil
}

// Cast casts an event along with its argument to the actor
func (a *actor) Cast(event Event, args ...Arg) error {
	a.RLock()
	defer a.RUnlock()

	if _, ok := a.eventMap[event]; !ok {
		return fmt.Errorf("eventEntity %v hasn't registered yet", event)
	}

	a.eventCh <- eventEntity{event: event, args: args}
	return nil
}

// Close closes the actor
func (a *actor) Close() error {
	a.closeOnce.Do(func() {
		close(a.quit)
		<-a.done
	})

	return nil
}

// process events one by one
func (a *actor) process() {
	defer close(a.done)

	for {
		select {
		case <-a.quit:
			return
		case evt := <-a.eventCh:
			h := a.eventMap[evt.event]
			h(evt.args...)
		}
	}
}
