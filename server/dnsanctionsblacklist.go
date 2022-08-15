// DN Sanctions List addresses

package server

import "strings"

var dnSanctionsBlacklist = map[string]bool{
	"0x0203298C575bf7efCeFB58C0ddd0aE7417439C19": true,
}

func isOnDNSanctionsList(address string) bool {
	addrs := strings.ToLower(address)
	return dnSanctionsBlacklist[addrs]
}
