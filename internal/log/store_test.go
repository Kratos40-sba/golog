package log

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenWidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := ioutil.TempFile("", "store_append_read_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)
	require.NoError(t, err)
	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)
	s, err = newStore(f)
	require.NoError(t, err)
}
func testAppend(t *testing.T, s *store) {

}
func testRead(t *testing.T, s *store) {

}
func testReadAt(t *testing.T, s *store) {

}
func TestStoreClose(t *testing.T) {

}
func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(
		name,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, 0, err
	}
	fileInformation, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}
	return f, fileInformation.Size(), nil
}
