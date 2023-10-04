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

	for _, restriction := range strings.Split(restrictions, ",") {
		RestrictTo[strings.TrimSpace(restriction)] = true
	}
}

// IsOperationAllowed checks if a given operation is allowed based on the --restrict-to flag.
func IsOperationAllowed(operation string) bool {
	allowed, exists := RestrictTo[operation]
	return exists && allowed
}
