package netease

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	*http.Client
}

var DefaultClient *Client

func init() {
	DefaultClient = new(Client)
	DefaultClient.Client = http.DefaultClient
}

func (client *Client) wrapRequest(req *http.Request) {
	req.Header.Add("Referer", "http://music.163.com/")
}

func (client *Client) Do(req *http.Request) ([]byte, error) {
	client.wrapRequest(req)
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (client *Client) Download(uri string, w io.Writer) (n int64, err error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return 0, err
	}
	client.wrapRequest(req)
	resp, err := client.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return io.Copy(w, resp.Body)
}

func GetSong(songId int64) ([]byte, error) {
	uri := fmt.Sprintf("http://music.163.com/api/song/detail?id=%d&ids=%%5B%d%%5D", songId, songId)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}

func GetArtist(artistId int64, limit int64) ([]byte, error) {
	uri := fmt.Sprintf("http://music.163.com/api/artist/albums/%d?limit=%d", artistId, limit)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}

func GetAlbum(albumId int64) ([]byte, error) {
	uri := fmt.Sprintf("http://music.163.com/api/album/%d", albumId)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}

func GetLyric(songId int64) ([]byte, error) {
	uri := fmt.Sprintf("http://music.163.com/api/song/media?id=%d", songId)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}

func GetTrack(dfsId int64) (s string) {
	var salt []byte = []byte("3go8&$8*3*3h0k(2)2")
	length := len(salt)
	var dfs []byte = []byte(strconv.FormatInt(dfsId, 10))
	for idx, b := range dfs {
		dfs[idx] = b ^ salt[idx%length]
	}
	var b = md5.Sum(dfs)
	var t = base64.StdEncoding.EncodeToString(b[0:])
	return fmt.Sprintf("http://m1.music.126.net/%s/%d.mp3", strings.Replace(strings.Replace(t, "/", "_", -1), "+", "-", -1), dfsId)
}

func GetPlayList(playListID int64) ([]byte, error) {
	uri := fmt.Sprintf("http://music.163.com/api/playlist/detail?id=%d", playListID)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	return DefaultClient.Do(req)
}
