// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gtag providing tag content storing for struct.
//
// Note that calling functions of this package is not concurrently safe,
// which means you cannot call them in runtime but in boot procedure.
package gtag

import (
	"github.com/pkg/errors"
	"regexp"
)

var (
	data  = make(map[string]string)
	regex = regexp.MustCompile(`\{(.+?)\}`)
)

// Set sets tag content for specified name.
// Note that it panics if `name` already exists.
func Set(name, value string) {
	if _, ok := data[name]; ok {
		panic(errors.Errorf(`value for tag name "%s" already exists`, name))
	}
	data[name] = value
}

// SetOver performs as Set, but it overwrites the old value if `name` already exists.
func SetOver(name, value string) {
	data[name] = value
}

// Sets sets multiple tag content by map.
func Sets(m map[string]string) {
	for k, v := range m {
		Set(k, v)
	}
}

// SetsOver performs as Sets, but it overwrites the old value if `name` already exists.
func SetsOver(m map[string]string) {
	for k, v := range m {
		SetOver(k, v)
	}
}

// Get retrieves and returns the stored tag content for specified name.
func Get(name string) string {
	return data[name]
}

// Parse parses and returns the content by replacing all tag name variable to
// its content for given `content`.
// Eg:
// gtag.Set("demo", "content")
// Parse(`This is {demo}`) -> `This is content`.
func Parse(content string) string {
	return regex.ReplaceAllStringFunc(content, func(s string) string {
		if v, ok := data[s[1:len(s)-1]]; ok {
			return v
		}
		return s
	})
}
