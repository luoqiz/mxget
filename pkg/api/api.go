package api

import (
	"context"

	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/utils"
)

type Song struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Artist    string `json:"artist"`
	Album     string `json:"album"`
	PicURL    string `json:"pic_url,omitempty"`
	Lyric     string `json:"lyric,omitempty"`
	ListenURL string `json:"listen_url,omitempty"`
}

type Rank struct {
	Sourceid string `json:"sourceid"`
	Intro    string `json:"intro"`
	Name     string `json:"name"`
	ID       string `json:"id"`
	Source   string `json:"source"`
	Pic      string `json:"pic"`
	Pub      string `json:"pub"`
}

type Collection struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	PicURL string  `json:"pic_url"`
	Songs  []*Song `json:"songs"`
}

type Provider interface {
	SearchSongs(ctx context.Context, keyword string, page int, pageSize int) ([]*Song, error)
	GetRank(ctx context.Context) ([]*Rank, error)
	GetSong(ctx context.Context, songId string) (*Song, error)
	GetArtist(ctx context.Context, artistId string, page int, pageSize int) (*Collection, error)
	GetAlbum(ctx context.Context, albumId string) (*Collection, error)
	GetPlaylist(ctx context.Context, playlistId string) (*Collection, error)
	SendRequest(req *ghttp.Request) (*ghttp.Response, error)
}

func (s *Song) String() string {
	return utils.PrettyJSON(s)
}

func (c *Rank) String() string {
	return utils.PrettyJSON(c)
}

func (c *Collection) String() string {
	return utils.PrettyJSON(c)
}
