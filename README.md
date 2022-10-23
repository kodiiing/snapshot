# Snapshot

Simple utility to do snapshot testing in Go.

```go
package main_test

import (
    "testing"

    "github.com/kodiiing/snapshot"
)

func TestSnapshot(t *testing.T) {
    var output string = SomeHeavyTask()

    ok, err := snapshot.MatchSnapshot("just-for-fun", output, snapshot.Config{})
    if err != nil {
        var snapshotError snapshot.SnapshotError
        if errors.As(err, &snapshotError) {
            t.Errorf(
                "Mismatched snapshot:\nDifferences: %d\nExpected: %s\nGot: %s",
                snapshotError.Difference,
                snapshotError.Snapshot,
                snapshotError.Received,
            )
        }
    }
}
```

For further documentation, see [pkg.go.dev](https://pkg.go.dev/github.com/kodiiing/snapshot)

## License

[MIT](./LICENSE)