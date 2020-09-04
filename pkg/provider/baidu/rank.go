package baidu

import (
	"context"
	"errors"
	"github.com/winterssy/mxget/pkg/api"
)

func (a *API) GetRank(ctx context.Context) ([]*api.Rank, error) {
	resp, err := a.SearchSongsRaw(ctx, "444", 1, 100)
	if err != nil {
		return nil, err
	}

	n := len(resp.Result.SongInfo.SongList)
	if n == 0 {
		return nil, errors.New("search songs: no data")
	}

	songs := make([]*api.Rank, n)
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
