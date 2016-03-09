package main

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/simulatedsimian/assert"
)

func TestFileIndex(t *testing.T) {
	fi := MakeFileIndex("data")

	assert.Equal(t, fi.makeFileName(0), "data/00000/000")
	assert.Equal(t, fi.makeFileName(999), "data/00000/999")
	assert.Equal(t, fi.makeFileName(1000), "data/00001/000")
	assert.Equal(t, fi.makeFileName(12345678), "data/12345/678")
}

func exist(fname string) bool {
	_, err := os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func TestStore(t *testing.T) {

	assert := assert.Make(t)

	path := filepath.Join("/tmp", "testdata", strconv.FormatInt(time.Now().UnixNano(), 10))

	fi := MakeFileIndex(path)

	err := fi.MakeDummyFiles(3999)
	assert(err).IsNil()

	fi.RefeshFileCount()
	assert(fi.fileCount).Equal(uint64(3999))

	assert(exist(filepath.Join(path, "00003", "999"))).Equal(false)
	index := fi.ReserveFileIndex()
	assert(fi.StoreFile(index, "test file")).NoError()
	assert(exist(filepath.Join(path, "00003", "999"))).Equal(true)

	assert(exist(filepath.Join(path, "00004", "000"))).Equal(false)
	index = fi.ReserveFileIndex()
	assert(fi.StoreFile(index, "test file")).NoError()
	assert(exist(filepath.Join(path, "00004", "000"))).Equal(true)
}
