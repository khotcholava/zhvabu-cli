package generator

import (
	"fmt"
	"os"
	"strings"
)

func GenerateComponentContent(componentName string, props []string, style string, memo bool, forwardRef bool, class bool) string {
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

	// Handle class component
	if class {
		componentFunc = generateClassComponent(componentName, props)
	} else {
		// Generate functional component with wrappers
		componentFunc = generateFunctionalComponent(componentName, props, memo, forwardRef)
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

// generateClassComponent generates a React class component
func generateClassComponent(componentName string, props []string) string {
	var propsType string
	if len(props) == 0 {
		propsType = "{}"
	} else {
		propsType = fmt.Sprintf("%sProps", componentName)
	}

	return fmt.Sprintf(`export class %s extends React.Component<%s, {}> {
  render() {
    return (
      <p>%s Component Works</p>
    );
  }
}`, componentName, propsType, componentName)
}

// generateFunctionalComponent generates a functional component with optional wrappers
func generateFunctionalComponent(componentName string, props []string, memo bool, forwardRef bool) string {
	var propsType string
	if len(props) == 0 {
		propsType = "{}"
	} else {
		propsType = fmt.Sprintf("%sProps", componentName)
	}

	// Base component body
	componentBody := `<p>%s Component Works</p>`

	// Handle forwardRef (must be applied first if both memo and forwardRef)
	if forwardRef {
		if memo {
			// Both memo and forwardRef
			return fmt.Sprintf(`export const %s = React.memo(React.forwardRef<%s, %s>((props, ref) => {
  return (
    %s
  );
}));`, componentName, propsType, propsType, fmt.Sprintf(componentBody, componentName))
		} else {
			// Only forwardRef
			return fmt.Sprintf(`export const %s = React.forwardRef<%s, %s>((props, ref) => {
  return (
    %s
  );
});`, componentName, propsType, propsType, fmt.Sprintf(componentBody, componentName))
		}
	} else if memo {
		// Only memo, no forwardRef
		return fmt.Sprintf(`export const %s = React.memo((props: %s) => {
  return (
    %s
  );
});`, componentName, propsType, fmt.Sprintf(componentBody, componentName))
	} else {
		// No wrappers, plain functional component
		return fmt.Sprintf(`export const %s = (props: %s) => {
  return (
    %s
  );
};`, componentName, propsType, fmt.Sprintf(componentBody, componentName))
	}
}
