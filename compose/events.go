package compose

type EventType byte

const (
	OnClick EventType = 10
)

func (e *EventType) IsEvent(et EventType) bool {
	return *e == et
}

type EventData struct {
	Type string
	data map[string]any
}

func NewEventData(t string) *EventData {
	return &EventData{Type: t, data: make(map[string]any)}
}
func (ed *EventData) GetValue(key string) any {
	if ed.data == nil {
		return nil
	}
	return ed.data[key]
}

func (ed *EventData) Setalue(key string, v any) {
	if ed.data == nil {
		return
	}
	ed.data[key] = v
}

type EventHandle func(sender Component, data *EventData)
