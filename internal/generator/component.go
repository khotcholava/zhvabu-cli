package generator

import (
	"fmt"
	"os"
	"strings"
)

func GenerateComponentContent(componentName string, props []string, style string) string {
	// Determine style extension
	styleExt := "css"
	if style != "" && style != "none" {
		styleExt = style
	}

	// CSS import (only if style is not "none")
	var cssImport string
	if style != "none" {
		cssImport = fmt.Sprintf("import './%s.%s';\n\n", componentName, styleExt)
	}

	reactImport := "import React from 'react';\n\n"
	var propsType string

	if len(props) == 0 {
		propsType = "" // No type when no props
	} else {
		propsList := make([]string, len(props))
		for i, prop := range props {
			propsList[i] = fmt.Sprintf("  %s: any", prop)
		}
		propsStr := strings.Join(propsList, ";\n")
		propsType = fmt.Sprintf("type %sProps = {\n%s\n};\n\n", componentName, propsStr)
	}

	var componentFunc string
	if len(props) == 0 {
		componentFunc = fmt.Sprintf(`export const %s = (props: {}) => {
  return (
    <p>%s Component Works</p>
  );
};`, componentName, componentName)
	} else {
		componentFunc = fmt.Sprintf(`export const %s = (props: %sProps) => {
  return (
    <p>%s Component Works</p>
  );
};`, componentName, componentName, componentName)
	}

	return cssImport + reactImport + propsType + componentFunc
}

func CreateStyleFile(componentName string, style string, path string) error {
	filename := fmt.Sprintf("%s/%s.%s", path, componentName, style)
	return os.WriteFile(filename, []byte(""), 0644)
}

func CreateIndexFile(componentName string, path string) error {
	content := fmt.Sprintf("export { %s } from \"./%s\";\n", componentName, componentName)
	return os.WriteFile(fmt.Sprintf("%s/index.ts", path), []byte(content), 0644)
}
