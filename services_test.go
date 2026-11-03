package neutrino

import "testing"

// TestServicesAdvertiseBlake2b checks the client says it follows the BLAKE2b
// rules, without which Bitcoin Knots from 29.4.3 sends it no headers past the
// fork.
func TestServicesAdvertiseBlake2b(t *testing.T) {
	if Services&SFNodeBlake2b == 0 {
		t.Fatalf("Services %v does not include SFNodeBlake2b", Services)
	}
}
