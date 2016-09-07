package dvb

import (
	"errors"
	"regexp"
	"strconv"
)

// A TransportMode e.g tram, citybus, etc.
type TransportMode struct {
	Title   string
	Name    string
	IconURL string
}

func parseMode(modeStr string) (mode *TransportMode, err error) {
	if intMode, convErr := strconv.Atoi(modeStr); convErr == nil {
		switch {
		case 0 <= intMode && intMode < 60:
			return TransportModes()["tram"], convErr
		case 60 <= intMode && intMode < 100:
			return TransportModes()["citybus"], convErr
		case 100 <= intMode && intMode < 1000:
			return TransportModes()["regiobus"], convErr
		}
	}

	if modeStr == "SWB" {
		return TransportModes()["lift"], nil
	}

	eR, _ := regexp.Compile("^E(\\d+)")
	if eR.Match([]byte(modeStr)) {
		matches := eR.FindStringSubmatch(modeStr)
		if intMode, convErr := strconv.Atoi(matches[1]); convErr == nil {
			switch {
			case 0 <= intMode && intMode < 60:
				return TransportModes()["tram"], convErr
			case 60 <= intMode && intMode < 100:
				return TransportModes()["citybus"], convErr
			}
		}
	}

	regioR, _ := regexp.Compile("^\\D$|^\\D\\/\\D$")
	if regioR.Match([]byte(modeStr)) {
		return TransportModes()["regiobus"], nil
	}

	ferryR, _ := regexp.Compile("^F")
	if ferryR.Match([]byte(modeStr)) {
		return TransportModes()["ferry"], nil
	}

	trainR, _ := regexp.Compile("^RE|^IC|^TL|^RB|^SB|^SE|^U\\d")
	if trainR.Match([]byte(modeStr)) {
		return TransportModes()["train"], nil
	}

	metroR, _ := regexp.Compile("^S")
	if metroR.Match([]byte(modeStr)) {
		return TransportModes()["metropolitan"], nil
	}

	astR, _ := regexp.Compile("alita")
	if astR.Match([]byte(modeStr)) {
		return TransportModes()["ast"], nil
	}

	err = errors.New("failed to parse departure identifier into transport mode")
	return
}

// TransportModes returns all possible modes of transport
func TransportModes() map[string]*TransportMode {
	return map[string]*TransportMode{
		"tram": {
			Title:   "Straßenbahn",
			Name:    "tram",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-tram.svg",
		},
		"citybus": {
			Title:   "Stadtbus",
			Name:    "citybus",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-citybus.svg",
		},
		"regiobus": {
			Title:   "Regionalbus",
			Name:    "regiobus",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-bus.svg",
		},
		"metropolitan": {
			Title:   "S-Bahn",
			Name:    "metropolitan",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-metropolitan.svg",
		},
		"lift": {
			Title:   "Seil-/Schwebebahn",
			Name:    "lift",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-lift.svg",
		},
		"ferry": {
			Title:   "Fähre",
			Name:    "ferry",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-ferry.svg",
		},
		"ast": {
			Title:   "Anrufsammeltaxi (AST)/ Rufbus",
			Name:    "ast",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-alita.svg",
		},
		"train": {
			Title:   "Zug",
			Name:    "train",
			IconURL: "https://www.dvb.de/assets/img/trans-icon/transport-train.svg",
		},
	}
}
