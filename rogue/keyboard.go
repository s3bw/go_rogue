package rogue

import "github.com/nsf/termbox-go"

type keyboardEventsType int

// Keyboard events
const (
	MOVE keyboardEventsType = 1 + iota
	RETRY
	END
)

type keyboardEvent struct {
	eventType keyboardEventsType
	key       termbox.Key
}

func keyToDirection(k termbox.Key) direction {
	switch k {
	case termbox.KeyArrowLeft:
		return LEFT
	case termbox.KeyArrowDown:
		return DOWN
	case termbox.KeyArrowRight:
		return RIGHT
	case termbox.KeyArrowUp:
		return UP
	default:
		return 0
	}
}

func listenToKeyboard(evChan chan keyboardEvent) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				evChan <- keyboardEvent{eventType: END, key: ev.Key}
			default:
				if ev.Ch == 'r' {
					evChan <- keyboardEvent{eventType: RETRY, key: ev.Key}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}