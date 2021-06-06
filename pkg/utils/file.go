package utils

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// File struct
type File struct {
	Name         string
	Size         int64
	FriendlySize string
	ModifiedTime time.Time
	IsDir        bool
	FullPath     string
	Ext          string
}

func IsValidPath(baseDir, filename string) (string, bool) {
	absPath, err := filepath.Abs(filepath.Join(baseDir, filename))
	if err != nil {
		return "", false
	}
	if strings.HasPrefix(absPath, baseDir) {
		return absPath, true
	}
	return "", false
}

// IsFileExists checks the file availability
func IsFileExists(filename string) *File {
	file, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return nil
	}

	f := getFile(file)
	return &f
}

// IsDirExists checks the directory availability
func IsDirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// ListFiles func
func ListFiles(baseDir, dir string) ([]File, error) {
	if !IsDirExists(dir) {
		return nil, errors.New("directory not found")
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	items := make([]File, 0)
	for _, file := range files {
		f := getFile(file)
		f.FullPath = strings.TrimPrefix(path.Join(dir, file.Name()), baseDir)
		items = append(items, f)
	}
	return items, nil
}

func getFile(file fs.FileInfo) File {
	fileSize := GetSize(file.Size())
	if file.IsDir() {
		fileSize = "-"
	}
	return File{
		Name:         file.Name(),
		Size:         file.Size(),
		FriendlySize: fileSize,
		ModifiedTime: file.ModTime(),
		IsDir:        file.IsDir(),
		Ext:          filepath.Ext(file.Name()),
	}
}
