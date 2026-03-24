//go:build race_demo

// To run the example, run:
//
//	go test -race -tags race_demo -run Example_copyRace ./internal/

package examples

// Example_copyRace demonstrates the potential benefit of the "nocopy" tag.
// Without it, users may encounter data-races.
func Example_copyRace() {
	var (
		s    OneOf // doesn't use nocopy
		done = make(chan struct{})
	)
	go func() {
		defer close(done)
		s.CalculateCanotoCache()
	}()

	// Read the atomic values concurrently with the writes.
	_ = s

	<-done
	// Output:
}
