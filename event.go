package tactics

type Action int

const (
	attackA Action = iota
	damagedA
	beforeCastA
	afterCastA
	timeGoA
	timeLineA // 无事件，标记为只在特定时间段生效
	shieldedA
	unShieldedA
	healedA
	healthPercentA
)

type Handler interface {
	handle(event Event, ground *Ground)
}

type Event struct {
	Action
	num int
}

func NewE(action Action, num int) Event {
	return Event{action, num}
}

func void(action Action) Event {
	return Event{action, 0}
}

func (e Event) Is(action Action) bool {
	return e.Action == action
}
