package dvb

import "testing"

type depPair struct {
	initAttrs   []string
	description string
}

var departureTests = []depPair{
	{[]string{"3", "Wilder Mann", ""}, "3 Wilder Mann in 0 minutes"},
	{[]string{"85", "Löbtau Süd", "5"}, "85 Löbtau Süd in 5 minutes"},
}

func TestInitDeparture(t *testing.T) {
	for _, pair := range departureTests {
		dep, err := initDeparture(pair.initAttrs)
		if err != nil {
			t.Error(err)
		}
		if dep.String() != pair.description {
			t.Error("Expected", pair.description, "got", dep.String())
		}
	}
}

func TestDepartureString(t *testing.T) {
	dep := &Departure{
		Line:         "3",
		Direction:    "Wilder Mann",
		RelativeTime: 5,
	}
	expected := "3 Wilder Mann in 5 minutes"
	if dep.String() != expected {
		t.Error("Expected", expected, "got", dep.String())
	}
}
