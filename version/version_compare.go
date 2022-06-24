package version

import (
	"strconv"
	"strings"
)

//VersionCompare
//compare version (x.x.x)
//
//a>b return 1
//
//a==b return 0
//
//a<b  return -1
func VersionCompare(a string, b string) int {
	aVersion := strings.Split(a, ".")
	bVersion := strings.Split(b, ".")
	for i, _ := range aVersion {
		aV, _ := strconv.Atoi(aVersion[i])
		bV, _ := strconv.Atoi(bVersion[i])

		if aV > bV {
			return 1
		} else if aV < bV {
			return -1
		}
	}
	return 0
}
