package xzoption

func WithName[T any](name string) Option[T] {
	return func(value *T) {
		SetValue[string](value, "Name", name)
	}
}
