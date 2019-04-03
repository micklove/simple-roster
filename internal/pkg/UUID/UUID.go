package UUID

type Generator interface {
	Create() (id string, err error)
}
