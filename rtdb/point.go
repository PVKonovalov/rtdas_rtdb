package rtdb

import "time"

type Point struct {
	Id           uint64 // Identifier
	Value        float64
	Source       uint16
	Qds          Quality // Quality descriptor
	Timestamp    IsoTime // Timestamp from RTU
	TimestampFep IsoTime // Timestamp from FEP
}

// Reserved Source IDs
const (
	SourceIdManualEntered uint16 = 1
	SourceIdCalculator    uint16 = 1
)

type Source struct {
	Id       uint64 // Identifier
	Priority uint16
	Point    Point
}

type Sources struct {
	sources []Source
}

// Reserved limit type identifiers
const (
	LimitTypeUpperWarn  uint16 = 1
	LimitTypeLowerWarn  uint16 = 2
	LimitTypeUpperAlarm uint16 = 3
	LimitTypeLowerAlarm uint16 = 4
)

type Limit struct {
	Id         uint64
	Value      float64
	TypeId     uint16
	TimeoutSec time.Duration // Timeout after which an alarm is generated
}

type Limits struct {
	limits []Limit
}

type RtdbPoint struct {
	Point    Point
	DeadZone float64
	TypeId   uint16
	Sources  Sources // List of data sources for the point
	Limits   Limits  // List of limits for the point
}
