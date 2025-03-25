package database

type IConnection interface {
	CreateConnection() error
}
