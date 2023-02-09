package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		environment string = "development"
	)

	enumFlag(&environment, "environment", []string{"development", "staging", "production"}, "Operating environment")

	flag.Parse()

	fmt.Printf("The operating environment is: %s\n", environment)
}

func enumFlag(target *string, name string, safelist []string, usage string) {
	flag.Func(name, usage, func(flagValue string) error {
		for _, allowedValue := range safelist {
			if flagValue == allowedValue {
				*target = flagValue
				return nil
			}
		}

		return fmt.Errorf("must be one of %v", safelist)
	})
}
