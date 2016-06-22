package checks

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"strings"
	"testing"
)

func InterestingGoroutines() (gs []string) {
	buf := make([]byte, 2<<20)
	buf = buf[:runtime.Stack(buf, true)]
	for _, g := range strings.Split(string(buf), "\n\n") {
		sl := strings.SplitN(g, "\n", 2)
		if len(sl) != 2 {
			continue
		}
		stack := strings.TrimSpace(sl[1])
		isNotInteresting :=
		stack == "" ||
		strings.Contains(stack, "runtime.goexit") ||
		strings.Contains(stack, ".InterestingGoroutines") ||
		strings.Contains(stack, "behavior_asynclooplogger") ||
		strings.Contains(stack, "resolver/resolver") ||
		strings.Contains(stack, "accumulator/accumulator") ||
		strings.Contains(stack, "os/signal.loop()") ||
		strings.Contains(stack, "bounce.go") ||
		strings.Contains(stack, "(*Clipboard).bundleCountersForAccumulator")

		if isNotInteresting {
			continue
		}
		gs = append(gs, stack)
	}
	sort.Strings(gs)
	return
}

// Verify the other tests didn't leave any goroutines running.
func GoroutinesLeaked() bool {
	if testing.Short() {
		// not counting goroutines for leakage in -short mode
		return false
	}
	gs := InterestingGoroutines()

	n := 0
	stackCount := make(map[string]int)
	for _, g := range gs {
		stackCount[g]++
		n++
	}

	if n == 0 {
		return false
	}
	fmt.Fprintf(os.Stderr, "Too many goroutines running after test(s).\n")
	for stack, count := range stackCount {
		fmt.Fprintf(os.Stderr, "%d instances of:\n%s\n", count, stack)
	}
	return true
}



