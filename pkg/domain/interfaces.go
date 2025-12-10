package domain

type HTTPStatusGetter interface {
	Error() string
	Status() int
}
