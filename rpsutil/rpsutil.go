// Package rpsutil provides utilities for building and configuring generic types through the use of option functions.
// It defines a `Lister` interface and a generic `Build` function for assembling a type with customizable options.
package rpsutil

import "reflect"

// Lister is a generic interface that represents a type that provides a list of configuration functions for type T.
// Each configuration function takes a pointer to T and applies specific settings to it.
type Lister[T any] interface {
	// List returns a slice of functions that each modify a given instance of type T.
	// These functions can be used to apply custom configurations to an instance of T.
	List() []func(*T) error
}

// Build creates a new instance of type T and applies all configuration functions provided by Lister options.
// It iterates over each option in opts and applies the contained functions to the new instance of T.
// If any configuration function returns an error, Build immediately returns nil and the encountered error.
//
// Parameters:
//   - opts: Variadic list of Lister implementations for type T, each containing a list of functions that modify T.
//
// Returns:
//   - *T: A pointer to the configured instance of type T.
//   - error: An error if any configuration function fails; otherwise, nil.
//
// Example usage:
//
//	type Config struct { /* fields */ }
//	config, err := rpsutil.Build(configOption1, configOption2)
//	if err != nil { /* handle error */ }
func Build[T any](opts ...Lister[T]) (*T, error) {

	t := new(T)

	for _, opt := range opts {
		if opt == nil || reflect.ValueOf(opt).IsNil() {
			continue
		}

		for _, setArgs := range opt.List() {

			if setArgs == nil {
				continue
			}

			if err := setArgs(t); err != nil {
				return nil, err
			}

		}

	}

	return t, nil
}
