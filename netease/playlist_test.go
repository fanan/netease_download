package netease

import (
	"log"
	"testing"
)

func TestPlayListIntergrate(t *testing.T) {
	var id int64 = 22914865
	pl := NewPlayList(id)
	err := pl.Parse()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(pl)
}
