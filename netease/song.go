package netease

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Track struct {
	DfsID     int64  `json:"dfsId"`
	Name      string `json:"name"`
	Extension string `json:"extension"`
}

type Song struct {
	//Name string `json:"name"`
	ID   int64  `json:"id"`
	Best *Track `json:"bMusic"`
}

type SongInfo struct {
	Songs []*Song `json:"songs"`
	Code  int     `json:"code"`
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	uri   string
}

func NewSongInfo(id int64) *SongInfo {
	return &SongInfo{ID: id, Songs: make([]*Song, 0, 1)}
}

func (si *SongInfo) Parse() (err error) {
	c, err := GetSong(si.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(c, si)
	if err != nil {
		return err
	}
	if si.Code != 200 {
		return fmt.Errorf("error code %d", si.Code)
	}
	if len(si.Songs) != 1 {
		return fmt.Errorf("error songs length: %d", len(si.Songs))
	}
	si.uri = GetTrack(si.Songs[0].Best.DfsID)
	return nil
}

func (si *SongInfo) Download(dir string) (err error) {
	if dir == "" {
		dir = os.TempDir()
	}
	var filename = strings.Replace(fmt.Sprintf("%s.%s", si.Songs[0].Best.Name, si.Songs[0].Best.Extension), "/", "-", -1)
	fp, err := os.Create(path.Join(dir, filename))
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = DefaultClient.Download(si.uri, fp)
	if err != nil {
		return err
	}
	return nil
}
