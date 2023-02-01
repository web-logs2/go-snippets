package main

import (
	"flag"
	"log"
	"strings"
	"time"
)

func main() {
	// First we need to declare variables to hold the values from the
	// command-line flags. Notice that we also need to set any defaults,
	// which will be used if the relevant flag is not provided at runtime.
	var (
		urls  []string                    // Default of the empty slice
		pause time.Duration = time.Second // Default of one second
	)

	// The flag.Func() function takes three parameters: the flag name,
	// descriptive help text, and a function with the signature
	// `func(string) error` which is called to process the string value
	// from the command-line flag at runtime and assign it to the necessary
	// variable. In this case, we use strings.Fields() to split the string
	// based on whitespace and store the resulting slice in the urls
	// variable that we declared above. We then return nil from the
	// function to indicate that the flag was parsed without any errors.
	flag.Func("urls", "List of URLs to print", func(flagValue string) error {
		urls = strings.Fields(flagValue)
		return nil
	})

	// Likewise we can do the same thing to parse duration. The
	// time.ParseDuration() function may throw an error here, so we make
	// sure to return that from our function.
	flag.Func("pause", "Duration to pause between printing URLs", func(flagValue string) error {
		var err error
		pause, err = time.ParseDuration(flagValue)
		return err
	})

	// Importantly, call flag.Value() to trigger actual parsing of the
	// flags.
	flag.Parse()

	// Print out the URLs, pausing between each iteration.
	for _, u := range urls {
		log.Print(u)
		time.Sleep(pause)
	}
}