package snapshot

import "fmt"

// SnapshotError contains difference between a snapshot and the received input.
type SnapshotError struct {
	Difference uint64
	Snapshot   string
	Received   string
}

func (s SnapshotError) Error() string {
	return fmt.Sprintf("mismatched snapshot with %d differences", s.Difference)
}
