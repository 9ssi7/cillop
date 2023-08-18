package server

type Server interface {
	Listen() error
}
