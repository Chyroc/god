package god

import (
	"strings"
)

var Str = &str{}

type str struct {
}

func (r *str) HasPrefixList(s string, prefixList ...string) bool {
	for _, prefix := range prefixList {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

func (r *str) CountPre(s, substr string) int {
	if s == "" || substr == "" {
		return 0
	}
	if strings.HasPrefix(s, substr) {
		return r.CountPre(s[len(substr):], substr) + 1
	}
	return 0
}

func (r *str) FindLastSubstr(s, substr string) string {
	if substr == "" {
		return ""
	}
	v := strings.Split(s, substr)
	if len(v) == 0 {
		return ""
	}
	return v[len(v)-1]
}
