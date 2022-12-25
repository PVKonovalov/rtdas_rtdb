package rtdb

import (
	"salyut_rtdb/qds"
	"strings"
	"time"
)

type IsoDate struct {
	time.Time
}

func (c *IsoDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`) // remove quotes
	c.Time, err = time.Parse("2006-01-02T15:04:05.999-0700", s)
	if err != nil {
		c.Time = time.Now()
	}
	return
}

func (c *IsoDate) MarshalJSON() ([]byte, error) {
	str := c.Time.Format("2006-01-02T15:04:05.999-0700")
	return []byte("\"" + str + "\""), nil
}

func (c IsoDate) String() string {
	return c.Time.Format("02.01.2006 15:04:05.999-0700")
}

const RetroPeriodSec = 5
const RetroLengthSec = 3600 * 24

type Point struct {
	id            uint64
	timestamp     IsoDate
	timestampRecv IsoDate
	value         float32
	quality       qds.Quality
	source        uint32
}

type Tag struct {
	name       string
	retro      [RetroLengthSec / RetroPeriodSec]Point
	retroIndex uint32
	typeId     uint32
}

func (tag *Tag) GetPoint() Point {
	return tag.retro[tag.retroIndex]
}

type Rtdb struct {
	points map[uint64]Tag
}
