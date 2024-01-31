package rtdb

import "sync"

type Rtdb struct {
	Points map[uint64]RtdbPoint
	sync.RWMutex
}
