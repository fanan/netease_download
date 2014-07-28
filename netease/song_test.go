package netease

import (
	"testing"
)

func TestSongIntegrated(t *testing.T) {
	var id int64 = 21398313
	si := NewSongInfo(id)
	err := si.Parse()
	if err != nil {
		t.Fatal(err)
	}
	dir := "."
	err = si.Download(dir)
	if err != nil {
		t.Fatal(err)
	}
}
