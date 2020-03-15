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
func (fm *FileManager) FindByRegex(rg string) []string {
	var files []string
	filepath.Walk(fm.path, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(rg, f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})

	return files
}

// DeleteByRegex delete files matched regexpr if cannot remove file
// its leave without error. Return teleted files count
func (fm *FileManager) DeleteByRegex(rg string) (deleted uint) {
	files := fm.FindByRegex(rg)
	if len(files) == 0 {
		return
	}

	for _, file := range files {
		err := os.Remove(filepath.Join(fm.path, file))
		if err == nil {
			deleted++
		}
	}

	return
}
