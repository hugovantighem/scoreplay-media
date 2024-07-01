package infra

func ToPtr[T any](item T) *T {
	return &item
}
