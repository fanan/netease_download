package netease

import (
	"encoding/json"
	"fmt"
)

type playListResult struct {
	ID     int64       `json:"id"`
	Name   string      `json:"name"`
	Tracks []*SongInfo `json:"tracks"`
}

type PlayList struct {
	Result *playListResult `json:"result"`
	Code   int             `json:"code"`
	id     int64
}

func NewPlayList(id int64) *PlayList {
	return &PlayList{id: id}
}

func (pl *PlayList) Parse() error {
	c, err := GetPlayList(pl.id)
	if err != nil {
		return err
	}
	err = json.Unmarshal(c, pl)
	if err != nil {
		return err
	}
	if pl.Code != 200 {
		return fmt.Errorf("error code: %d", pl.Code)
	}
	return nil
}

func (pl *PlayList) String() string {
	var s string = fmt.Sprintf("%s\n", pl.Result.Name)
	for idx, track := range pl.Result.Tracks {
		s += fmt.Sprintf("%d: %s: %d\n", idx, track.Name, track.ID)
	}
	return s
}
