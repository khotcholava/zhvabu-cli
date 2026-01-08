package parser

import (
	"fmt"
	"strings"
)

func ParseProps(propsStr string) ([]string, error) {
	if propsStr == "" {
		return []string{}, nil
	}

	props := strings.Split(propsStr, ",")
	var filteredProps []string
	for _, prop := range props {
		trimmed := strings.TrimSpace(prop)

		if trimmed != "" {
			if err := ValidatePropName(trimmed); err != nil {
				return nil, err
			}
			filteredProps = append(filteredProps, trimmed)
		}
	}

	return filteredProps, nil
}

func ValidatePropName(name string) error {
	if name == "" {
		return fmt.Errorf("empty prop name detected (trailing comma?)")
	}

	firstChar := name[0]

	// Check if starts with letter or underscore
	isLetter := (firstChar >= 'a' && firstChar <= 'z') || (firstChar >= 'A' && firstChar <= 'Z')
	isUnderscore := firstChar == '_'

	if !isLetter && !isUnderscore {
		return fmt.Errorf("invalid prop name: '%s' (must start with letter or underscore)", name)
	}

	// Check if starts with number (invalid)
	if firstChar >= '0' && firstChar <= '9' {
		return fmt.Errorf("invalid prop name: '%s' (cannot start with number)", name)
	}

	// Check for spaces
	if strings.Contains(name, " ") {
		return fmt.Errorf("invalid prop name: '%s' (contains spaces)", name)
	}

	// Check for hyphens
	if strings.Contains(name, "-") {
		return fmt.Errorf("invalid prop name: '%s' (contains hyphen, use camelCase)", name)
	}

	// Check for dots
	if strings.Contains(name, ".") {
		return fmt.Errorf("invalid prop name: '%s' (contains dot)", name)
	}

	// Check if all chars are valid

	for _, char := range name {
		isLetter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
		isNumber := char >= '0' && char <= '9'
		isUnderscore := char == '_'

		if !isLetter && !isNumber && !isUnderscore {
			return fmt.Errorf("invalid prop name: '%s' (contains invalid character: '%c')", name, char)
		}
	}

	return nil
}
