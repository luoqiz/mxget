package tencent

import (
	"context"
	"errors"
	"fmt"
	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/api"
	"math/rand"
	"strings"
)

func (a *API) GetPlayLists(ctx context.Context, page int, pageSize int) ([]*api.Playlist, error) {
	resp, err := a.GetPlaylistsRaw(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	n := len(resp.Data.List)
	if n == 0 {
		return nil, errors.New("search songs: no data")
	}

	playlists := make([]*api.Playlist, n)
	for i, s := range resp.Data.List {
		playlists[i] = &api.Playlist{
			ID:        s.Dissid,
			Name:      strings.TrimSpace(s.DissName),
			Img:       strings.TrimSpace(s.Imgurl),
			Total:     "0",
			ListenNum: s.Listennum,
		}
	}
	return playlists, nil
}

// 获取歌单列表
func (a *API) GetPlaylistsRaw(ctx context.Context, page int, pageSize int) (*PlaylistsResponse, error) {
	params := ghttp.Params{
		"rnd": rand.Float32(),
		"sin": (page - 1) * pageSize,
		"ein": ((page - 1) * pageSize) + 19,
	}

	resp := new(PlaylistsResponse)
	req, _ := ghttp.NewRequest(ghttp.MethodGet, apiGetPlaylists)
	req.SetQuery(params)
	req.SetContext(ctx)
	r, err := a.SendRequest(req)
	if err == nil {
		err = r.JSON(resp)
	}
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, fmt.Errorf("get playlist: %d", resp.Code)
	}

	return resp, nil
}
