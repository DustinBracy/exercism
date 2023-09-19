package gigasecond

import "time"

// AddGigasecond adds 1,000,000,000 seconds to the input time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
