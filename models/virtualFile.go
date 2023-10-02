package models

// VirtualFile represents a file or directory in the virtual file system.
type VirtualFile struct {
	IsDir      bool
	DirectLink string
}
