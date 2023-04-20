package contracts

type EncryptDriver func(string) Encryptor

type EncryptManager interface {

	// Extend 通过给定的key和加密器扩展加密工厂
	// extend encryption factory with given key and encryptor.
	Extend(driver string, provider EncryptDriver)

	// Encryptor 通过给定的名称获取加密器
	// Get the encryptor by the given key.
	Encryptor(name string) Encryptor

	// Driver 通过给定的名称获取加密驱动
	// Get the driver by the given key.
	Driver(driver string) EncryptDriver
}

type Encryptor interface {
	// Encrypt 加密给定的值
	// Encrypt the given value.
	Encrypt(value []byte) []byte
	EncryptString(value string) string

	// Decrypt 解密给定的值
	// Decrypt the given value.
	Decrypt(encrypted []byte) ([]byte, error)
	DecryptString(encrypted string) (string, error)
}
