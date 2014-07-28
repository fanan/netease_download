package main

import (
	"flag"
	"github.com/fanan/netease_download/netease"
	"log"
	"os"
)

var playlistID = flag.Int64("l", 22914865, "playlistID")
var downloadDir = flag.String("d", os.ExpandEnv("$HOME/Downloads/"), "download dir")

func main() {
	flag.Parse()
	var pl = netease.NewPlayList(*playlistID)
	fi, err := os.Lstat(*downloadDir)
	if err != nil {
		log.Fatal(err)
	}
	if !fi.IsDir() {
		log.Fatalf("%s is not a directory", *downloadDir)
	}
	err = pl.Parse()
	if err != nil {
		log.Fatal(err)
	}
	n := len(pl.Result.Tracks)
	for idx, track := range pl.Result.Tracks {
		log.Printf("start downloading %d/%d -- %s\n", idx+1, n, track.Name)
		err = track.Parse()
		if err != nil {
			log.Fatal(err)
		}
		err = track.Download(*downloadDir)
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Downloading finished!")
}
