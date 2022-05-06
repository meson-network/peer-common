package api

import (
	"errors"

	"github.com/coreservice-io/dns-common/spec00"
)

func ResolveNodeDomainFromIp(ip string) (string, error) {
	prefix, err := spec00.GenerateSpecStr(ip)
	if err == nil {
		return prefix + "." + DOMAIN, nil
	} else {
		return "", errors.New("resolveDomain error")
	}
}

// func ResolveNodeIpFromDomain(domain string) (string, error) {

// 	return "", nil
// }
