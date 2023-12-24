package storage

type IRepository interface {
	DoRequest(request string) error
}
