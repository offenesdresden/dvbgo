package dvb

import (
	"fmt"
	"strconv"
)

// Departure encapsulates info regarding the Line, Direction and relative departure time in minutes.
type Departure struct {
	Line         string
	Direction    string
	RelativeTime int
}

func (dep *Departure) String() string {
	return fmt.Sprintf("%s %s in %d minutes", dep.Line, dep.Direction, dep.RelativeTime)
}

func initDeparture(attrs []string) (departure *Departure, err error) {
	rel, err := strconv.Atoi(attrs[2])
	if err != nil {
		return
	}
	departure = &Departure{
		Line:         attrs[0],
		Direction:    attrs[1],
		RelativeTime: rel,
	}
	return
}
