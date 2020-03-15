package fm

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// Files used to manipuliate with file list
type Files struct {
	list []string
}

// NewFiles create and return new Files
func NewFiles() *Files {
	return &Files{
		list: make([]string, 0),
	}
}

// Append add file name to list
func (f *Files) Append(name string) {
	if name == "" {
		return
	}

	f.list = append(f.list, name)
}

// Remove used to remove file name from list
func (f *Files) Remove(name string) (ok bool) {
	for i, file := range f.list {
		if file == name {
			f.list = append(f.list[:i], f.list[i+1:]...)
			ok = true
			break
		}
	}

	return
}

// Count return files count in Files list
func (f *Files) Count() int {
	return len(f.list)
}

// DeleteAll delete all files in path name
func (f *Files) DeleteAll(path string) (deleted int) {
	for _, file := range f.list {
		err := os.Remove(filepath.Join(path, file))
		if err == nil {
			deleted++
		}
	}
	return
}

// FindLastVersion search and return file name with last version
func (f *Files) FindLastVersion(r string) string {
	rg := regexp.MustCompile(r)

	var maxMajor, maxMinor, maxPatch, maxBuild int
	var major, minor, patch, build int
	var name string

	for _, file := range f.list {
		vs := rg.FindAllString(file, -1)

		switch len(vs) {
		case 1:
			major, _ = strconv.Atoi(vs[0])
			if maxMajor > major {
				continue
			}
		case 2:
			major, _ = strconv.Atoi(vs[0])
			minor, _ = strconv.Atoi(vs[1])
			if maxMajor > major || maxMinor > minor {
				continue
			}
		case 3:
			major, _ = strconv.Atoi(vs[0])
			minor, _ = strconv.Atoi(vs[1])
			patch, _ = strconv.Atoi(vs[2])
			if maxMajor > major || maxMinor > minor || maxPatch > patch {
				continue
			}
		case 4:
			major, _ = strconv.Atoi(vs[0])
			minor, _ = strconv.Atoi(vs[1])
			patch, _ = strconv.Atoi(vs[2])
			build, _ = strconv.Atoi(vs[3])
			if maxMajor > major || maxMinor > minor || maxPatch > patch || maxBuild > build {
				continue
			}
		default:
			continue
		}

		maxMajor = major
		maxMinor = minor
		maxPatch = patch
		maxBuild = build

		name = file
	}

	return name
}
