package xiami

import (
	"context"
	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/api"
)

func (a *API) GetRank(ctx context.Context) ([]*api.Rank, error) {
	//params := ghttp.Params{
	//	"musicId": mid,
	//}

	//resp := new(SongLyricResponse)
	req, _ := ghttp.NewRequest(ghttp.MethodGet, "")
	//req.SetQuery(params)
	req.SetContext(ctx)
	r, err := a.SendRequest(req)
	println(r)
	if err == nil {
		//err = r.JSON(resp)
	}
	if err != nil {
		return nil, err
	}

	//if resp.Status != 200 {
	//	return nil, fmt.Errorf("get song lyric: %s", resp.Msg)
	//}

	//
	//resp, err := a.SearchSongsRaw(ctx, keyword, page, pageSize)
	//if err != nil {
	//	return nil, err
	//}
	//
	//n := len(resp.Result.SongInfo.SongList)
	//if n == 0 {
	//	return nil, errors.New("search songs: no data")
	//}

	songs := make([]*api.Rank, 1)
	//for i, s := range resp.Result.SongInfo.SongList {
	//	songs[i] = &api.Song{
	//		Id:     s.SongId,
	//		Name:   strings.TrimSpace(s.Title),
	//		Artist: strings.TrimSpace(strings.ReplaceAll(s.Author, ",", "/")),
	//		Album:  strings.TrimSpace(s.AlbumTitle),
	//	}
	//}
	return songs, nil
}
