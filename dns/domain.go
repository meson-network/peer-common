package dns

import (
	"errors"

	"github.com/coreservice-io/dns-common/spec00"
)

func ResolveNodeDomainFromIp(ip string, root_domain string) (string, error) {
	prefix, err := spec00.GenerateSpecStr(ip)
	if err == nil {
		return prefix + "." + root_domain, nil
	} else {
		return "", errors.New("resolveDomain error")
	}
}
