package fsutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "fsutil")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	from := filepath.Join(dir, "from.txt")
	to := filepath.Join(dir, "to.txt")

	f, err := os.Create(from)
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("Hello, World")
	f.Close()

	err = CopyFile(to, from)
	if err != nil {
		t.Fatal(err)
	}

	ff, err := os.Stat(from)
	if err != nil {
		t.Fatal(err)
	}

	tf, err := os.Stat(to)
	if err != nil {
		t.Fatal(err)
	}

	if ff.Size() != tf.Size() {
		t.Fatalf("the file both should have same size: from %v to %v", ff.Size(), tf.Size())
	}

	if ff.Mode() != tf.Mode() {
		t.Fatalf("the file both should have same mode: from %v to %v", ff.Mode(), tf.Mode())
	}
}
