package gigasecond

import "time"
import "math"

// AddGigasecond adds 100^9  seconds to the input time t
func AddGigasecond(t time.Time) time.Time {
	secondsToAdd := int(math.Pow(10, 9))
	return t.Add(time.Second * time.Duration(secondsToAdd))
}
