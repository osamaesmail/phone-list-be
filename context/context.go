package context

type Context interface {
	BindQuery(obj interface{}) error
}
