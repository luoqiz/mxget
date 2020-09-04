package kuwo

import (
	"context"
	"fmt"
	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/api"
	"strings"
)

func (a *API) GetRank(ctx context.Context) ([]*api.Rank, error) {

	resp := new(RankResponse)
	req, _ := ghttp.NewRequest(ghttp.MethodGet, apiGetRank)
	req.SetContext(ctx)
	r, err := a.SendRequest(req)

	if err == nil {
		err = r.JSON(resp)
	}
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, fmt.Errorf("get rank : %s", resp.Msg)
	}
	total := 0
	for _, rankInfo := range resp.Data {
		total = total + len(rankInfo.List)
	}

	ranks := make([]*api.Rank, total)
	index := 0
	for _, rankInfo := range resp.Data {
		for _, rank := range rankInfo.List {
			ranks[index] = &api.Rank{
				Sourceid: strings.TrimSpace(rank.Sourceid),
				Intro:    strings.TrimSpace(rank.Intro),
				ID:       strings.TrimSpace(rank.ID),
				Source:   strings.TrimSpace(rank.Source),
				Name:     strings.TrimSpace(rank.Name),
				Pub:      strings.TrimSpace(rank.Pub),
				Pic:      strings.TrimSpace(rank.Pic),
			}
			index++
		}

	}
	return ranks, nil
}
