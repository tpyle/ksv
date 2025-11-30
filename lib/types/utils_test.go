package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixList(t *testing.T) {
	cases := []struct {
		prefix string
		list   []string
		expect []string
	}{
		{">prefix", []string{"a", "b", "c"}, []string{">prefix>a", ">prefix>b", ">prefix>c"}},
		{">prefix", []string{""}, []string{">prefix"}},
	}
	for _, c := range cases {
		result := PrefixList(c.prefix, c.list)
		assert.Equal(t, c.expect, result)
	}
}

func TestPrefixMap(t *testing.T) {
	cases := []struct {
		prefix string
		m      map[string]string
		expect map[string]string
	}{
		{">prefix", map[string]string{"a": "1", "b": "2"}, map[string]string{">prefix>a": "1", ">prefix>b": "2"}},
		{">prefix", map[string]string{"": "1"}, map[string]string{">prefix": "1"}},
	}
	for _, c := range cases {
		result := PrefixMap(c.prefix, c.m)
		assert.Equal(t, c.expect, result)
	}
}

func TestUnprefixMap(t *testing.T) {
	cases := []struct {
		prefix string
		m      map[string]string
		expect map[string]string
	}{
		{">prefix", map[string]string{">prefix/a": "1", ">prefix/b": "2"}, map[string]string{"a": "1", "b": "2"}},
	}
	for _, c := range cases {
		result := UnprefixMap(c.prefix, c.m)
		assert.Equal(t, c.expect, result)
	}
}

func TestSplitPath(t *testing.T) {
	cases := []struct {
		path   string
		expect []string
	}{
		{">a>b>c", []string{"", "a", "b", "c"}},
		{"a>b>c", []string{"a", "b", "c"}},
	}
	for _, c := range cases {
		result := SplitPath(c.path)
		assert.Equal(t, c.expect, result)
	}
}
