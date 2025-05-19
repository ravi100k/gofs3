// backend/interface.go
package backend

type StorageBackend interface {
	ListObjects() ([]string, error)
}
