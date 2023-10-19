package address

import "strings"

// Verify if address is valid or not
func AddressType(address string) string {
	validTypes := []string{"rua", "avenida", "estrada", "rodovia"}

	firstWordAddress := strings.ToLower(strings.Split(address, " ")[0])

	for _, word := range validTypes {
		if word == firstWordAddress {
			return strings.Title(firstWordAddress)
		}
	}
	return "Invalid Address"
}
