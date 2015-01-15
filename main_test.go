package ziputil

import (
	"io"
	"os"
	"testing"
)

func TestReadTopLevelFile(t *testing.T) {
	zipFile, err := os.Open("testFiles/625ab.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer zipFile.Close()

	aReader, err := FileFromZipReader(zipFile, "625a.txt")
	if err != nil {
		t.Fatal(err)
	}
	aReader.Close()

	buf := make([]byte, 625)

	n, err := aReader.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 625 {
		t.Fatal("expected to read 625 bytes")
	}

	for i, b := range buf {
		if b != byte('a') {
			t.Fatalf("Expected every byte to be an 'a' but found %q at position %d", b, i)
		}
	}

	n, err = aReader.Read(buf)
	if err != io.EOF {
		t.Log("expected no more content")
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("expected no more content")
	}
}

func TestReadNestedFile(t *testing.T) {
	zipFile, err := os.Open("testFiles/625ab_nested.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer zipFile.Close()

	nestedBReader, err := FileFromZipReader(zipFile, "nested/625b.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer nestedBReader.Close()

	buf := make([]byte, 625)

	n, err := nestedBReader.Read(buf)
	if err != nil {
		t.Fatal(err)
	}

	if n != 625 {
		t.Fatal("expected to read 625 bytes")
	}

	for i, b := range buf {
		if b != byte('b') {
			t.Fatalf("Expected every byte to be an 'b' but found %q at position %d", b, i)
		}
	}

	n, err = nestedBReader.Read(buf)
	if err != io.EOF {
		t.Log("expected no more content")
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatal("expected no more content")
	}
}
