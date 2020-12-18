package kuwo

import (
	"context"
	"errors"
	"fmt"
	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/api"
	"strings"
)

func (a *API) GetPlayLists(ctx context.Context, page int, pageSize int) ([]*api.Playlist, error) {
	resp, err := a.PlayListSongsRaw(ctx, page, pageSize)
	if err != nil {
		return nil, err
	}

	n := len(resp.Data.Data)
	if n == 0 {
		return nil, errors.New("search songs: no data")
	}

	playlists := make([]*api.Playlist, n)
	for i, s := range resp.Data.Data {
		playlists[i] = &api.Playlist{
			ID:        s.ID,
			Name:      strings.TrimSpace(s.Name),
			Img:       strings.TrimSpace(s.Img),
			Total:     s.Total,
			ListenNum: s.Listencnt / 10000,
		}
	}
	return playlists, nil
}

// 获取歌单榜单
func (a *API) PlayListSongsRaw(ctx context.Context, page int, pageSize int) (*PlayListsResponse, error) {
	params := ghttp.Params{
		"order": "new",
		"pn":    page,
		"rn":    pageSize,
	}

	resp := new(PlayListsResponse)
	req, _ := ghttp.NewRequest(ghttp.MethodGet, apiGetPlayLists)
	req.SetQuery(params)
	req.SetContext(ctx)
	r, err := a.SendRequest(req)
	if err == nil {
		err = r.JSON(resp)
	}
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		if resp.Code == -1 {
			err = errors.New("search songs: no data")
		} else {
			err = fmt.Errorf("search songs: %s", resp.errorMessage())
		}
		return nil, err
	}

	return resp, nil
}
