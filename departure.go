package dvb

import (
	"fmt"
	"strconv"
)

type Departure struct {
	Line string
	Direction string
	RelativeTime int
}

func (dep *Departure) Mode() {
	// TODO: Implement me
	panic("Not yet implemented")
}

func (dep *Departure) String() string {
	return fmt.Sprintf("%s %s in %d minutes", dep.Line, dep.Direction, dep.RelativeTime)
}

func InitDeparture(attrs []string) (*Departure, error) {
	rel, err := strconv.Atoi(attrs[2])
	if err != nil {
		return nil, err
	}
	return &Departure{
		Line: attrs[0],
		Direction: attrs[1],
		RelativeTime: rel,
	}, nil
}
