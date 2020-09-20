package kuwo

import (
	"context"
	"errors"
	"fmt"
	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/api"
	"strconv"
	"strings"
)

func (a *API) GetRankList(ctx context.Context, bangId string, page int, pageSize int) ([]*api.Song, error) {
	resp, err := a.GetRankListRaw(ctx, bangId, page, pageSize)
	if err != nil {
		return nil, err
	}

	n := len(resp.Data.MusicList)
	if n == 0 {
		return nil, errors.New("rank list songs: no data")
	}

	songs := make([]*api.Song, n)
	for i, s := range resp.Data.MusicList {
		songs[i] = &api.Song{
			Id:       strconv.Itoa(s.Rid),
			Name:     strings.TrimSpace(s.Name),
			Artist:   strings.TrimSpace(strings.ReplaceAll(s.Artist, "、", "/")),
			ArtistId: strconv.Itoa(s.Artistid),
			Album:    strings.TrimSpace(s.Album),
			AlbumId:  strconv.Itoa(s.Albumid),
			AlbumPic: strings.TrimSpace(s.Albumpic),
			Duration: s.Duration,
		}
	}
	return songs, nil
}

// 获取榜单里的歌曲
func (a *API) GetRankListRaw(ctx context.Context, bangId string, page int, pageSize int) (*RankListResponse, error) {
	rankId, _ := strconv.Atoi(bangId)
	params := ghttp.Params{
		"bangId": rankId,
		"pn":     page,
		"rn":     pageSize,
	}

	fmt.Println(params)

	resp := new(RankListResponse)
	req, _ := ghttp.NewRequest(ghttp.MethodGet, apiGetRankListInfo)
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
		return nil, fmt.Errorf("rank list songs: %s", resp.errorMessage())
	}
	return resp, nil
}
