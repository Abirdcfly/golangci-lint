//golangcitest:args -Egocritic
//golangcitest:config_path testdata/configs/gocritic-fix.yml
//golangcitest:expected_exitcode 1
package p

import (
	"strings"
)

func gocritic() {
	var xs [2048]byte

	// xs -> &xs
	for _, x := range xs {
		print(x)
	}

	// strings.Count("foo", "bar") == 0 -> !strings.Contains("foo", "bar")
	if strings.Count("foo", "bar") == 0 {
		print("qu")
	}
}
