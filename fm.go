package fm

import (
	"os"
	"path/filepath"
	"regexp"
)

// FileManager used for erase old unused dll file from bin directory
type FileManager struct {
	path string
}

// NewFM validate path and if it valid, then return new FileManager
func NewFM(path string) (*FileManager, error) {
	_, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	return &FileManager{
		path: path,
	}, nil
}

// FindByRegex find files in path and return list
func (fm *FileManager) FindByRegex(rg string) *Files {
	files := NewFiles()
	filepath.Walk(fm.path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(rg, f.Name())
			if err == nil && r {
				files.Append(f.Name())
			}
		}
		return nil
	})

	return files
}

// DeleteByRegex delete files matched regexpr if cannot remove file
// its leave without error. Return teleted files count
func (fm *FileManager) DeleteByRegex(rg string) int {
	files := fm.FindByRegex(rg)
	if files.Count() == 0 {
		return 0
	}

	return files.DeleteAll(fm.path)
}
