package dvb

import "testing"

type modePair struct {
	identifier string
	modeName   string
}

var modeTests = []modePair{
	{"3", "tram"},
	{"11", "tram"},
	{"59", "tram"},
	{"E8", "tram"},
	{"E11", "tram"},
	{"E", "tram"},
	{"85", "citybus"},
	{"99", "citybus"},
	{"E75", "citybus"},
	{"EV2", "citybus"},
	{"366", "regiobus"},
	{"999", "regiobus"},
	{"A", "regiobus"},
	{"Z", "regiobus"},
	{"G/L", "regiobus"},
	{"H/S", "regiobus"},
	{"SWB", "lift"},
	{"F7", "ferry"},
	{"F14", "ferry"},
	{"ICE 1717", "train"},
	{"IC 1818", "train"},
	{"RB 1919", "train"},
	{"TLX 2020", "train"},
	{"SB33", "train"}, // "Sächsische Städtebahn"
	{"SE19", "train"}, // "Wintersport Express" o.O
	{"U28", "train"},  // fares between Bad Schandau and Děčín
	{"S3", "metropolitan"},
	{"S 2121", "metropolitan"},
	{"alita", "ast"},
	{"alita 95", "ast"},
}

func TestMode(t *testing.T) {
	for _, pair := range modeTests {
		v, err := parseMode(pair.identifier)
		if err != nil {
			t.Error(err)
		}
		if v.Name != pair.modeName {
			t.Error("For", pair.identifier, "expected", pair.modeName, "got", v.Name)
		}
	}
}
