package fm

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewFM(t *testing.T) {
	fm, err := NewFM("/notValidPath")
	if err == nil {
		t.Errorf("Wanted error, got nil :(")
	}
	if fm != nil {
		t.Errorf("Wanted FileManager nil, got structure :(")
	}

	fm, err = NewFM(".")
	if err != nil {
		t.Errorf("Wanted error nil, got: [%s]", err)
	}

	want := &FileManager{path: "."}
	if !reflect.DeepEqual(fm, want) {
		t.Errorf("Wanted: %v\nGot: %v", want, fm)
	}
}

func TestFindByRegex(t *testing.T) {
	dir, err := ioutil.TempDir(".", "tfm")
	if err != nil {
		t.Errorf("Cannot create temp directory: %s", err)
	}
	defer os.RemoveAll(dir)

	to := 3
	for i := 0; i < to; i++ {
		_, err := ioutil.TempFile(dir, "tfm_*.tmp")
		if err != nil {
			t.Errorf("%s", err)
		}
	}

	fm, err := NewFM(dir)
	if err != nil {
		t.Errorf("Wanted error nil, got: [%s]", err)
	}

	files := fm.FindByRegex(`^tfm_\d*.tmp`)
	if files == nil {
		t.Errorf("Wanted files %d, got nil", to)
	}

	deleted := fm.DeleteByRegex(`^bfm_\d*.tmp`)
	if deleted != 0 {
		t.Errorf("Wanted zero deleted files, got %d", deleted)
	}

	deleted = fm.DeleteByRegex(`^tfm_\d*.tmp`)
	if deleted != to {
		t.Errorf("Wanted deleted files: %d, got %d", to, deleted)
	}

	deleted = fm.DeleteAllFiles(files)
	if deleted != 0 {
		t.Errorf("Wanted zero, go %d", deleted)
	}

	if fm.GetPath() != dir {
		t.Errorf("Wrong FileManager path, wanted: %s, got: %s", dir, fm.GetPath())
	}
}
