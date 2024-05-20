package rest

type APIResponse[T any] struct {
	StatusCode int
	Message    string
	Data       T
}
