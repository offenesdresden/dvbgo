package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	Monitor("Helmholtzstraße", 0, 10, "Dresden")
}

// Monitor returns monitoring information for a single stop.
// An optional offset offsets information minutes into the future.
// Use the limit param to limit the number of results. Unfortunately it's never
// guaranteed that you'll get as many results as requested.
// City defaults to Dresden, but others in the VVO are possible as well.
func Monitor(stop string, offset int, limit int, city string) {
	if limit == 0 {
		limit = 10
	}
	if city == "" {
		city = "Dresden"
	}

	vvoURL, _ := url.Parse("http://widgets.vvo-online.de/abfahrtsmonitor/Abfahrten.do")
	params := vvoURL.Query()
	params.Set("ort", city)
	params.Set("hst", stop)
	params.Set("vz", string(offset))
	params.Set("lim", string(limit))
	vvoURL.RawQuery = params.Encode()

	resp, err := http.Get(vvoURL.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Body)
}
