package store

type Store interface {
	Push(key, value string) error
	Pop(key string) (string, error)
	Enqueue(key, value string) error
	Dequeue(key string) (string, error)
	Append(key, value string) error
	Remove(key string) (string, error)
	GetState(key string) ([]string, error)
}

var Default Store

func Init(host, port string) {
	Default = NewRedisStore(host, port)
}

func Push(key, value string) error          { return Default.Push(key, value) }
func Pop(key string) (string, error)        { return Default.Pop(key) }
func Enqueue(key, value string) error       { return Default.Enqueue(key, value) }
func Dequeue(key string) (string, error)    { return Default.Dequeue(key) }
func Append(key, value string) error        { return Default.Append(key, value) }
func Remove(key string) (string, error)     { return Default.Remove(key) }
func GetState(key string) ([]string, error) { return Default.GetState(key) }
