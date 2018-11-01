package compare

import (
	"strings"
	"time"
)

// String compares string values.
func String(old, new string) bool {
	return strings.Compare(strings.Trim(old, " "), strings.Trim(new, " ")) == 0
}

// Int compares integer values.
func Int(old, new int) bool {
	return old == new
}

// Date compares time values.
func Date(old, new time.Time) bool {
	return (old.Year() == new.Year()) && (old.YearDay() == new.YearDay())
}

// URLVisited checks if url has already been visited.
func URLVisited(err error) bool {
	return strings.Contains(err.Error(), "URL already visited")
}
