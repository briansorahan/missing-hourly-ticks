package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		newline = byte(0x0A)
		prev    time.Time
		r       = bufio.NewReader(os.Stdin)
	)
ReadLoop:
	for {
		line, err := r.ReadString(newline)
		if err == io.EOF {
			break ReadLoop
		}
		if err != nil {
			panic(err)
		}
		line = strings.TrimSpace(line)

		tm, err := time.Parse(time.RFC3339, line)
		if err != nil {
			panic(err)
		}
		if prev.IsZero() {
			prev = tm
			continue ReadLoop
		}
		if tm.Sub(prev) > time.Hour {
			missing := getMissing(prev, tm)
			// fmt.Printf("%s - %s -> %s\n", tm.Format(time.RFC3339), prev.Format(time.RFC3339), tm.Sub(prev))
			// fmt.Printf("missing -> %s\n", missing)
			for _, tm := range missing {
				fmt.Printf("%s\n", tm.Format(time.RFC3339))
			}
		}
		prev = tm
	}
}

func getMissing(start, end time.Time) []time.Time {
	var ret []time.Time
	for tm := start.Add(time.Hour); tm.Before(end); tm = tm.Add(time.Hour) {
		ret = append(ret, tm)
		// fmt.Printf("%s\n", tm.Format(time.RFC3339))
	}
	return ret
}
