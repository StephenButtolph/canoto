//go:build race_demo

// To run the example, run:
//
//	go test -race -tags race_demo -run Example_copyRace ./internal/

package examples

// Example_copyRace demonstrates the potential benefit of the "nocopy" tag.
// Without it, users may encounter data-races.
func Example_copyRace() {
	var s OneOf
	go s.CalculateCanotoCache()
	_ = s

	// Output:
}
