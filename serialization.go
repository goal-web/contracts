package contracts

type Serialization interface {
	Method(name string) Serializer
	Extend(name string, serializer Serializer)
}

type Serializer interface {
	Serialize(interface{}) string
	Unserialize(string, interface{}) error
}
