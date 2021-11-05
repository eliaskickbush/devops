package cache

type IDao interface {
	SetValue(key string, value interface{}) error
	GetValue(key string) (string,error)
	ClearValue(key string) error
}