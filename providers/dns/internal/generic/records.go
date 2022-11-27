package generic

import "fmt"

// GetRecordName extracts the leading DNS label(s) from fqdn, relative
// to authZone (which might be empty).
//
// Notes:
// - fqdn and authZone are expected to both be fully qualified DNS names
// - authZone is expected to be the DNS name for the SOA record of fqdn
func GetRecordName(fqdn, authZone string) (string, error) {
	end := len(fqdn) - len(authZone) - 1

	if len(fqdn) < end || end < 0 {
		return "", fmt.Errorf("%d is lower than the length of the fqdn (fqdn: %s, authZone: %s)", end, fqdn, authZone)
	}

	return fqdn[0:end], nil
}
