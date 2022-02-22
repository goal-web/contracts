package contracts

// Validatable 可验证的表单
type Validatable interface {
	FieldsProvider

	Rules() Fields
}

type ShouldValidate interface {
	Validatable

	ShouldVerify() bool
}
