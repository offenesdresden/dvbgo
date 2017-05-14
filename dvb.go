package dvb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Monitor returns a list of upcoming departures at a given stop.
// Offset offsets information minutes into the future, 0 meaning now.
func Monitor(stop string, offset int, city string) (departures []*Departure, err error) {
	// city no longer defaults to Dresden
	// this breaks the DVB short names for stops
	// see https://github.com/kiliankoe/dvbgo/issues/6
	/*
	if city == "" {
		city = "Dresden"
	}
	*/

	// for no apparent reason, you can no longer search for stops by a single char
	// see https://github.com/kiliankoe/dvbgo/issues/5
	if len(stop) <= 1 {
		return
	}

	vvoURL, _ := url.Parse("http://widgets.vvo-online.de/abfahrtsmonitor/Abfahrten.do")
	params := vvoURL.Query()
	params.Set("ort", city)
	params.Set("hst", stop)
	params.Set("vz", strconv.Itoa(offset))
	// somehow this breaks the GET request, the response will always be []
	// see https://github.com/kiliankoe/dvbgo/issues/4
	// params.Set("lim", "0")
	vvoURL.RawQuery = params.Encode()

	resp, err := http.Get(vvoURL.String())
	if err != nil {
		return
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var departuresData [][]string
	err = json.Unmarshal(respData, &departuresData)
	if err != nil {
		return
	}

	for _, depAttrs := range departuresData {
		departure, _ := initDeparture(depAttrs)
		departures = append(departures, departure)
	}

	return
}
