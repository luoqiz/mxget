package netease

import (
	"context"
	"github.com/winterssy/mxget/pkg/api"
)

func (a *API) GetRankList(ctx context.Context, bangId string, page int, pageSize int) ([]*api.Song, error) {
	songs := make([]*api.Song, 1)
	return songs, nil
}
