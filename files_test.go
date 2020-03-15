package fm

import (
	"testing"
)

func TestFiles(t *testing.T) {
	f := NewFiles()
	f.Append("")
	if f.Count() != 0 {
		t.Errorf("Wanted 0, got %d", f.Count())
	}

	f.Append("tst_v1.0.0.tst")
	f.Append("tst_v2.2.19.tst")
	f.Append("tst_v2.2.20.tst")
	f.Append("tst_v2.1.19.tst")
	f.Append("tst_v2.0.0.tst")
	f.Append("tst_v1.0.19.tst")

	name := f.FindLastVersion(`\d+`)
	if name != "tst_v2.2.20.tst" {
		t.Errorf("Wanted v2.2.20 file, got %s", name)
	}

	if f.Remove("testing.txt") {
		t.Errorf("Wanted return false, got true")
	}

	if !f.Remove(name) {
		t.Errorf("Wanted return true, got false")
	}
}
