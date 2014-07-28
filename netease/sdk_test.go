package netease

import (
	"log"
	"os"
	"testing"
)

func TestGetSong(t *testing.T) {
	var id int64 = 4337372
	s, err := GetSong(id)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(s))
}

func TestGetArtist(t *testing.T) {
	var id int64 = 32671
	var limit int64 = 50
	s, err := GetArtist(id, limit)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(s))
}

func TestGetAlbum(t *testing.T) {
	var id int64 = 1615951
	s, err := GetAlbum(id)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(s))
}

func TestGetLyric(t *testing.T) {
	var id int64 = 4337372
	s, err := GetLyric(id)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(s))
}

func TestGetTrack(t *testing.T) {
	var id int64 = 2065982348600621
	s := GetTrack(id)
	if s != "http://m1.music.126.net/rTK_rH6gRtaSPUlsBTA_9A==/2065982348600621.mp3" {
		t.Logf("need http://m1.music.126.net/rTK_rH6gRtaSPUlsBTA_9A==/2065982348600621.mp3\n")
		t.Fatalf("got  %s\n", s)
	}
}

func TestDownloadTrack(t *testing.T) {
	t.SkipNow()
	var id int64 = 2046191139300605
	uri := GetTrack(id)
	log.Println(uri)
	log.Println("start downloading")
	fp, err := os.Create("Yesterday.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	n, err := DefaultClient.Download(uri, fp)
	if err != nil || n != 5029068 {
		t.Log(err)
		t.Logf("n=%d", n)
		t.Fatal()
	}
}

func TestGetPlayList(t *testing.T) {
	var pid int64 = 22914865
	c, err := GetPlayList(pid)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(c))
}
