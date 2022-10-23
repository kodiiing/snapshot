// snapshot provides simple utility to do snapshot testing
package snapshot

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

// Config defines the configuration of the snapshot test behavior.
// If the struct is left empty, it will use the default configuration.
type Config struct {
	// AlwaysUpdate instructs that the snapshot file will always be updated.
	AlwaysUpdate bool
	// SnapshotDirectory sets directory location for snapshot file to exists.
	// If not provided, it defaults to the current directory.
	SnapshotDirectory string
}

// MatchSnapshot will finds if a snapshot is different from the given input.
// If the file does not exists, it will be created.
// Title must be unique for all the existing snapshot test cases on your code base,
// as we can't guarantee the order consistency of your tests.
//
// Differences on snapshot can be checked with asserting the error to the
// SnapshotError struct.
func MatchSnapshot(title string, input string, config Config) (bool, error) {
	snapshotFileName := path.Join(config.SnapshotDirectory, title+".snap")
	file, err := os.Open(snapshotFileName)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return false, fmt.Errorf("snapshot: opening file - %w", err)
	}
	defer func() {
		if file == nil {
			return
		}

		err := file.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			log.Printf("[defer] snapshot: closing file: %s", err.Error())
			return
		}
	}()

	var fileContents io.Reader
	if errors.Is(err, os.ErrNotExist) || config.AlwaysUpdate {
		fileContents = bytes.NewBufferString(input)
		// Create the file
		err := os.WriteFile(snapshotFileName, []byte(input), 0644)
		if err != nil {
			return false, fmt.Errorf("snapshot: writing file - %w", err)
		}
	} else {
		fileContents = file
	}

	// Store output here
	var currentLineOfCode uint64
	var lastDifferenceFoundLineOfCode uint64
	var differences uint64
	var snapshotOutput strings.Builder
	var receivedOutput strings.Builder

	// Read line by line
	snapshotScanner := bufio.NewScanner(fileContents)
	snapshotScanner.Split(bufio.ScanLines)
	inputScanner := bufio.NewScanner(strings.NewReader(input))
	inputScanner.Split(bufio.ScanLines)

	for {
		// Iterate each scanner
		snapshotNext := snapshotScanner.Scan()
		inputNext := inputScanner.Scan()

		// Break if nothing could be scan
		if !snapshotNext && !inputNext {
			break
		}
		// Get each line
		currentLineOfCode += 1
		snapshotContent := snapshotScanner.Text()
		inputContent := inputScanner.Text()

		if snapshotContent == inputContent {
			continue
		}

		// Snapshot and input does not match
		differences += 1

		if lastDifferenceFoundLineOfCode > 0 && currentLineOfCode-lastDifferenceFoundLineOfCode > 1 {
			snapshotOutput.WriteString("...\n")
			receivedOutput.WriteString("...\n")
		}

		snapshotOutput.WriteString(strconv.FormatUint(currentLineOfCode, 10))
		snapshotOutput.WriteString(" | ")
		snapshotOutput.WriteString(snapshotContent)
		snapshotOutput.WriteString("\n")
		receivedOutput.WriteString(strconv.FormatUint(currentLineOfCode, 10))
		receivedOutput.WriteString(" | ")
		receivedOutput.WriteString(inputContent)
		receivedOutput.WriteString("\n")

		lastDifferenceFoundLineOfCode = currentLineOfCode
	}

	if differences == 0 {
		return true, nil
	}

	return false, SnapshotError{
		Snapshot:   strings.TrimSpace(snapshotOutput.String()),
		Received:   strings.TrimSpace(receivedOutput.String()),
		Difference: differences,
	}
}
