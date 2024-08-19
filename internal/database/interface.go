package database

//go:generate mockgen -source interface.go -destination mocks/mock_database.go -package database_mocks
type Database interface {
}
