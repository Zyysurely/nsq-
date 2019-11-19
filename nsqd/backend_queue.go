package nsqd

// BackendQueue represents the behavior for the secondary message
// storage system
// BackendQueue时nsq中的二级存储，存入磁盘的，DummyBackendQueue（临时的）和 diskqueue
type BackendQueue interface {
	Put([]byte) error
	ReadChan() chan []byte // this is expected to be an *unbuffered* channel
	Close() error
	Delete() error
	Depth() int64
	Empty() error
}
