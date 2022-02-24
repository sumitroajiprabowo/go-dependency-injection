package example

import "fmt"

type File struct {
	Name string
}

func (f *File) Open() {
	fmt.Println("Opening file:", f.Name)
}

func NewFile(name string) (*File, func()) {
	file := &File{
		Name: name,
	}
	return file, func() {
		file.Open()
	}
}

type FileSystem struct {
	*File
}

func (f *FileSystem) Open() {
	fmt.Println("Opening file from filesystem:", f.File.Name)
}

func NewFileSystem(file *File) (*FileSystem, func()) {
	fileSystem := &FileSystem{
		File: file,
	}
	return fileSystem, func() {
		fileSystem.Open()
	}
}
