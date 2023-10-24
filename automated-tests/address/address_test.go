package address

import (
	"testing"
)

// sempre começar com Test
func TestTypeAddress(t *testing.T) {
	addressForTest := "avenida"
	expectedAddressType := "Rua"
	receivedAddressType := AddressType(addressForTest)

	if expectedAddressType != receivedAddressType {
		t.Errorf("The received type is different from the expected parameter! Expected %s and received %s",
			expectedAddressType,
			receivedAddressType)
	}
}
