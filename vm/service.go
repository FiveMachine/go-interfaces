package vm

type Service interface {
	New(context Context) (Instance, error)
	Source() Source
	Close() error
}
