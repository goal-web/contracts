package contracts

type EncryptorFactory interface {
	Encryptor

	// Extend 通过给定的key和加密器扩展加密工厂
	// extend encryption factory with given key and encryptor.
	Extend(key string, encryptor Encryptor)

	// Driver 通过给定的key获取加密器
	// Get the encryptor by the given key.
	Driver(key string) Encryptor
}

type Encryptor interface {
	// Encode 加密给定的值
	// Encrypt the given value.
	Encode(value string) string

	// Decode 解密给定的值
	// Decrypt the given value.
	Decode(encrypted string) (string, error)
}
