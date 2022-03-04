package contracts

import (
	"bufio"
	"io/fs"
	"time"
)

type FileVisibility int

// FileSystemProvider 通过给定的名称和配置信息获取文件系统
// Get the filesystem with the given name and configuration information.
type FileSystemProvider func(name string, config Fields) FileSystem

type File interface {
	fs.FileInfo
	// Read 读取文件
	// read a file.
	Read() []byte

	// ReadString 读取文件并返回字符串
	// read a file and return a string.
	ReadString() string

	// Disk 获取文件系统名称
	// get filesystem name.
	Disk() string
}

type FileSystemFactory interface {
	// Disk 通过给定的名称获取文件系统实现
	// Get a filesystem implementation.
	Disk(disk string) FileSystem

	// Extend 通过给定的名称和文件系统提供者扩展文件系统工厂
	// Extends the filesystem factory with the given name and filesystem provider.
	Extend(driver string, provider FileSystemProvider)

	FileSystem
}

type FileSystem interface {
	// Name 从文件路径中提取文件名
	// Extract the file name from a file path.
	Name() string


	// Exists 确定文件或目录是否存在
	// Determine if a file or directory exists.
	Exists(path string) bool


	// Get 获取给定路径文件的内容
	// get the contents of a file.
	Get(path string) (string, error)

	// Read 读取给定路径文件
	// read a file.
	Read(path string) ([]byte, error)

	// ReadStream 检索路径的读取流
	// Retrieves a read-stream for a path.
	ReadStream(path string) (*bufio.Reader, error)


	// Put 创建文件或更新（如果存在）
	// Create a file or update if exists.
	Put(path, contents string) error

	// WriteStream 使用流写入新文件
	// Write a new file using a stream.
	WriteStream(path string, contents string) error


	// GetVisibility 获取文件的可见性
	// get a file's visibility.
	GetVisibility(path string) FileVisibility

	// SetVisibility 设置文件的可见性
	// Set the visibility for a file.
	SetVisibility(path string, perm fs.FileMode) error


	// Prepend 预置到文件中
	// prepend to a file.
	Prepend(path, contents string) error


	// Append 附加到文件
	// append to a file.
	Append(path, contents string) error


	// Delete 删除给定路径的文件
	// Delete the file at a given path.
	Delete(path string) error


	// Copy 将文件复制到新位置
	// copy a file to a new location.
	Copy(from, to string) error


	// Move 将文件移动到新位置
	// move a file to a new location.
	Move(from, to string) error


	// Size 获取给定文件的文件大小
	// get the file size of a given file.
	Size(path string) (int64, error)


	// LastModified 获取文件的最后修改时间
	// Get the file's last modification time.
	LastModified(path string) (time.Time, error)


	// Files 获取目录中所有文件的数组
	// get an array of all files in a directory.
	Files(directory string) []File


	// AllFiles 从给定目录中获取所有文件（递归）
	// get all of the files from the given directory (recursive).
	AllFiles(directory string) []File


	// Directories 获取给定目录中的所有目录
	// get all of the directories within a given directory.
	Directories(directory string) []string


	// AllDirectories 从给定目录中获取所有目录（递归）
	// get all directories from a given directory (recursive).
	AllDirectories(directory string) []string


	// MakeDirectory 创建一个目录
	// Create a directory.
	MakeDirectory(path string) error


	// DeleteDirectory 递归删除目录
	// Recursively delete a directory.
	DeleteDirectory(directory string) error
}
