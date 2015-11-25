package core

type Event struct {
	Type string
	Target *EventDispatcher
}

func NewEvent(type_str string) (*Event) {
	return &Event{
		Type: type_str,
	}
}

type Listener func(Event)

type EventDispatcher struct {
	listeners map[string]([]Listener)
}

func NewEventDispatcher() (*EventDispatcher) {
	return &EventDispatcher{
		listeners: make(map[string]([]Listener), 0),
	}
}

func (e *EventDispatcher) indexOfListener(listeners []Listener, thelistener func(Event)) int {
	for i, listener := range listeners {
		if listener == thelistener {
			return i
		}
	}
	return -1
}

func (e *EventDispatcher) AddEventListener(type_str string, listener func(Event)) {
	listeners := e.listeners

	if _, ok := listeners[type_str]; !ok {
		listeners[type_str] = make([]Event, 0)
	}

	if ( e.indexOfListener(listeners[type_str], listener ) == - 1 ) {
		listeners[type_str] = append(listeners[type_str], listener )
	}
}

func (e *EventDispatcher) HasEventListener(type_str string, listener Listener) bool {
	listeners := e.listeners
	if a, ok := listeners[type_str]; ok && e.indexOfListener(a, listener ) != - 1 {
		return true
	}
	return false
}

func (e *EventDispatcher) RemoveEventListener(type_str string, listener Listener) {
	listeners := e.listeners
	if _, ok := listeners[type_str]; ok {
		var index = e.indexOfListener(listeners[type_str], listener)
		if index != - 1 {
			listeners[type_str] = append(listeners[type_str][:index], listeners[type_str][index + 1:])
		}
	}
}

func (e *EventDispatcher) DispatchEvent(event Event) {
	listeners := e.listeners
	if listenerArray, ok := listeners[event.Type]; ok {
		event.Target = e

		var array = make([]Listener, len(listenerArray))
		var length = len(listenerArray)

		for i := 0; i < length; i ++ {
			array[ i ] = listenerArray[ i ];
		}

		for i := 0; i < length; i++ {
			array[i](event);
		}
	}
}
