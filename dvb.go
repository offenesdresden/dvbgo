package dvbgo

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

}
