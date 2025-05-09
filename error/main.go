package main

import (
	"errors"
	"fmt"
	"io"
)

// Define a sentinel error
var ErrNotFound = errors.New("item not found")

// Custom error type
type ValidationError struct {
	Field string
	Msg   string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", v.Field, v.Msg)
}

// A function that returns an error
func findItem(id int) (string, error) {
	if id == 0 {
		return "", ErrNotFound
	}
	if id < 0 {
		return "", &ValidationError{Field: "id", Msg: "must be positive"}
	}
	return "ItemFound", nil
}

// A function that wraps errors
func processItem(id int) error {
	item, err := findItem(id)
	if err != nil {
		return fmt.Errorf("processing item %d failed: %w", id, err)
	}
	fmt.Println("Processing:", item)
	return nil
}

func main() {
	// Test with different values
	ids := []int{1, 0, -2}

	for _, id := range ids {
		err := processItem(id)
		if err != nil {
			// Handle known sentinel error
			if errors.Is(err, ErrNotFound) {
				fmt.Println("Handled NotFound:", err)
			} else {
				// Handle custom error using errors.As
				var ve *ValidationError
				if errors.As(err, &ve) {
					fmt.Println("Handled ValidationError:", ve)
				} else {
					// Generic error handler
					fmt.Println("Unhandled error:", err)
				}
			}
		} else {
			fmt.Println("Item processed successfully.")
		}
	}

	// Unwrapping example
	originalErr := io.EOF
	wrappedErr := fmt.Errorf("read failed: %w", originalErr)
	fmt.Println("Wrapped Error:", wrappedErr)

	unwrapped := errors.Unwrap(wrappedErr)
	fmt.Println("Unwrapped Error:", unwrapped)
}
