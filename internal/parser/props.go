package parser

import "strings"

func ParseProps(propsStr string) []string {
	if propsStr == "" {
		return []string{}
	}
	props := strings.Split(propsStr, ",")
	for i, prop := range props {
		props[i] = strings.TrimSpace(prop)
	}
	return props
}
