package address

import (
	"testing"
)

// sempre começar com Test
func TestTypeAddress(t *testing.T) {
	addressForTest := "Rua xpto"
	expectedAddressType := "Rua"
	receivedAddressType := AddressType(addressForTest)

	if expectedAddressType != receivedAddressType {
		t.Error("The received type is different from the expected parameter")
	}
}
