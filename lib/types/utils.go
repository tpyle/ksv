package types

import (
	"strings"

	"github.com/sirupsen/logrus"
)

const separator = "/"

func PrefixList(prefix string, l []string) []string {
	prefixed := make([]string, len(l))
	for i, v := range l {
		if v == "" {
			if len(l) == 1 {
				prefixed[i] = prefix
			} else {
				logrus.Errorf("empty path element under prefix %s", prefix)
				continue
			}
		} else {
			prefixed[i] = prefix + separator + v
		}
	}
	return prefixed
}

func PrefixMap(prefix string, m map[string]string) map[string]string {
	prefixed := make(map[string]string)
	for k, v := range m {
		if k == "" {
			if len(m) == 1 {
				prefixed[prefix] = v
			} else {
				logrus.Errorf("empty path element under prefix %s", prefix)
				continue
			}
		} else {
			prefixed[prefix+separator+k] = v
		}
	}
	return prefixed
}

func UnprefixMap(prefix string, m map[string]string) map[string]string {
	unprefixed := make(map[string]string)
	for k, v := range m {
		if len(k) < len(prefix) {
			continue
		}
		if k[:len(prefix)] != prefix {
			continue
		}
		unprefixed[k[len(prefix)+1:]] = v
	}
	return unprefixed
}

func SplitPath(path string) []string {
	return strings.Split(path, separator)
}
