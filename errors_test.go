package snapshot_test

import (
	"errors"
	"testing"

	"github.com/kodiiing/snapshot"
)

func TestSnapshotError(t *testing.T) {
	var err error = snapshot.SnapshotError{Received: "Hello world", Snapshot: "Hello world!", Difference: 1}

	var snapshotError snapshot.SnapshotError
	if errors.As(err, &snapshotError) {
		if snapshotError.Difference != 1 {
			t.Errorf("expecting `snapshotError.Difference` to be 1, got %d", snapshotError.Difference)
		}

		if snapshotError.Received != "Hello world" {
			t.Errorf("expecting `snapshotError.Received` to be \"Hello world\", got %s", snapshotError.Received)
		}

		if snapshotError.Snapshot != "Hello world!" {
			t.Errorf("expecting `snapshotError.Snapshot` to be \"Hello world!\", got %s", snapshotError.Snapshot)
		}
	} else {
		t.Error("expecting `err` to be asserted as SnapshotError struct, failed")
	}
}

func TestSnapshotError_Error(t *testing.T) {
	var err error = snapshot.SnapshotError{Received: "Hello world", Snapshot: "Hello world!", Difference: 1}

	if err.Error() != "mismatched snapshot with 1 differences" {
		t.Errorf("expected `err.Error()` to be: %s, got %s", "mismatched snapshot with 1 differences", err.Error())
	}
}
