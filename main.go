package main

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb.SetKeys(keybd_event.VK_F20)

	for {
		pressKey(kb)
		time.Sleep(90 * time.Second)
	}
}

func pressKey(kb keybd_event.KeyBonding) {
	kb.Press()
	time.Sleep(10 * time.Millisecond)
	kb.Release()
}
