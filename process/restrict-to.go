// process/restrict-to.go

package process

import (
	"strings"
)

// RestrictTo holds the operations that are allowed based on the --restrict-to flag.
var RestrictTo map[string]bool

// ParseRestrictTo parses the --restrict-to flag and sets the allowed operations.
func ParseRestrictTo(restrictions string) {
	RestrictTo = make(map[string]bool)

	// Check if the restrictions string is not empty
	if restrictions == "" {
		return
	}

	for _, restriction := range strings.Split(restrictions, ",") {
		RestrictTo[strings.TrimSpace(restriction)] = true
	}
	// debug, print the restrictions
	// log.Printf("Restrictions: %v", RestrictTo)
	// log.Printf("Length of restrictions: %v", len(RestrictTo))
}

// IsOperationAllowed checks if a given operation is allowed based on the --restrict-to flag.
func IsOperationAllowed(operation string) bool {
	// If RestrictTo is empty, allow all operations
	if len(RestrictTo) == 0 {
		// log.Printf("No restrictions set, allowing all operations")
		return true
	}

	allowed, exists := RestrictTo[operation]
	return exists && allowed
}
