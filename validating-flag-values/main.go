package main

import (
	"errors"
	"flag"
	"fmt"
)

func main() {
	var (
		environment string = "development"
	)

	flag.Func("environment", "Operating environment", func(flagValue string) error {
		for _, allowedValue := range []string{"development", "staging", "production"} {
			if flagValue == allowedValue {
				environment = flagValue
				return nil
			}
		}

		return errors.New(`must be one of "development", "staging" or "production""`)
	})

	flag.Parse()

	fmt.Printf("The operating environment is: %s\n", environment)
}
