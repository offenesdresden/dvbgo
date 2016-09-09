package dvb

import (
	"testing"

	_ "github.com/orchestrate-io/dvr"
)

func TestMonitor(t *testing.T) {
	deps, err := Monitor("Postplatz", 0, "")
	if err != nil {
		t.Error(err)
	}
	if len(deps) == 0 {
		t.Error("Expected more than 0 departures from 'Postplatz'")
	}
}
