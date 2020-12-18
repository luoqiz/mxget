package kuwo

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/winterssy/ghttp"
	"github.com/winterssy/mxget/pkg/api"
	"github.com/winterssy/mxget/pkg/concurrency"
	"github.com/winterssy/mxget/pkg/request"
	"github.com/winterssy/mxget/pkg/utils"
)

const (
	apiSearch          = "http://www.kuwo.cn/api/www/search/searchMusicBykeyWord"
	apiGetSong         = "http://www.kuwo.cn/api/www/music/musicInfo"
	apiGetSongURL      = "http://www.kuwo.cn/url?format=mp3&response=url&type=convert_url3"
	apiGetSongLyric    = "http://www.kuwo.cn/newh5/singles/songinfoandlrc"
	apiGetArtistInfo   = "http://www.kuwo.cn/api/www/artist/artist"
	apiGetArtistSongs  = "http://www.kuwo.cn/api/www/artist/artistMusic"
	apiGetAlbum        = "http://www.kuwo.cn/api/www/album/albumInfo"
	apiGetPlaylist     = "http://www.kuwo.cn/api/www/playlist/playListInfo"
	apiGetRank         = "http://www.kuwo.cn/api/www/bang/bang/bangMenu"
	apiGetRankListInfo = "http://www.kuwo.cn/api/www/bang/bang/musicList"
	apiGetPlayLists    = "http://www.kuwo.cn/api/pc/classify/playlist/getRcmPlayList"
	songDefaultBR      = 320
)

var (
	std = New(request.DefaultClient)
)

type (
	CommonResponse struct {
		Code int    `json:"code"`
		Msg  string `json:"msg,omitempty"`
	}

	Song struct {
		RId             int    `json:"rid"`
		Name            string `json:"name"`
		ArtistId        int    `json:"artistid"`
		Artist          string `json:"artist"`
		AlbumId         int    `json:"albumid"`
		Album           string `json:"album"`
		AlbumPic        string `json:"albumpic"`
		Track           int    `json:"track"`
		IsListenFee     bool   `json:"isListenFee"`
		SongTimeMinutes string `json:"songTimeMinutes"`
		Lyric           string `json:"-"`
		URL             string `json:"-"`
	}

	Rank struct {
		Sourceid string `json:"sourceid"`
		Intro    string `json:"intro"`
		Name     string `json:"name"`
		ID       string `json:"id"`
		Source   string `json:"source"`
		Pic      string `json:"pic"`
		Pub      string `json:"pub"`
	}

	RankInfo struct {
		Name string `json:"name"`
		List []Rank `json:"list"`
	}

	RankResponse struct {
		CommonResponse
		CurTime   int64      `json:"curTime"`
		ProfileID string     `json:"profileId"`
		ReqID     string     `json:"reqId"`
		Data      []RankInfo `json:"data"`
	}

	RankMusicList struct {
		Musicrid         string   `json:"musicrid"`
		Artist           string   `json:"artist"`
		Mvpayinfo        struct{} `json:"-"`
		Trend            string   `json:"trend"`
		Pic              string   `json:"pic"`
		Isstar           int      `json:"isstar"`
		Rid              int      `json:"rid"`
		Duration         int      `json:"duration"`
		Score100         string   `json:"score100"`
		ContentType      string   `json:"content_type"`
		RankChange       string   `json:"rank_change"`
		Track            int      `json:"track"`
		HasLossless      bool     `json:"hasLossless"`
		Hasmv            int      `json:"hasmv"`
		ReleaseDate      string   `json:"releaseDate"`
		Album            string   `json:"album"`
		Albumid          int      `json:"albumid"`
		Pay              string   `json:"pay"`
		Artistid         int      `json:"artistid"`
		Albumpic         string   `json:"albumpic"`
		Originalsongtype int      `json:"originalsongtype"`
		SongTimeMinutes  string   `json:"songTimeMinutes"`
		IsListenFee      bool     `json:"isListenFee"`
		Pic120           string   `json:"pic120"`
		Name             string   `json:"name"`
		Online           int      `json:"online"`
		PayInfo          struct{} `json:"-"`
		Nationid         string   `json:"nationid,omitempty"`
	}

	RankListData struct {
		Img       string          `json:"img"`
		Num       string          `json:"num"`
		Pub       string          `json:"pub"`
		MusicList []RankMusicList `json:"musicList,omitempty"`
	}

	RankListResponse struct {
		CommonResponse
		CurTime   int64        `json:"curTime"`
		ProfileID string       `json:"profileId"`
		ReqID     string       `json:"reqId"`
		Data      RankListData `json:"data"`
	}

	SearchSongsResponse struct {
		CommonResponse
		Data struct {
			Total string  `json:"total"`
			List  []*Song `json:"list"`
		} `json:"data"`
	}

	SongResponse struct {
		CommonResponse
		Data Song `json:"data"`
	}

	SongURLResponse struct {
		CommonResponse
		URL string `json:"url"`
	}

	SongLyricResponse struct {
		Status int    `json:"status"`
		Msg    string `json:"msg,omitempty"`
		Data   struct {
			LrcList []struct {
				Time      string `json:"time"`
				LineLyric string `json:"lineLyric"`
			} `json:"lrclist"`
		} `json:"data"`
	}

	ArtistInfo struct {
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Pic300 string `json:"pic300"`
	}

	ArtistInfoResponse struct {
		CommonResponse
		Data ArtistInfo `json:"data"`
	}

	ArtistSongsResponse struct {
		CommonResponse
		Data struct {
			List []*Song `json:"list"`
		} `json:"data"`
	}

	AlbumResponse struct {
		CommonResponse
		Data struct {
			AlbumId   int     `json:"albumId"`
			Album     string  `json:"album"`
			Pic       string  `json:"pic"`
			MusicList []*Song `json:"musicList"`
		} `json:"data"`
	}

	PlaylistResponse struct {
		CommonResponse
		Data struct {
			Id        int     `json:"id"`
			Name      string  `json:"name"`
			Img700    string  `json:"img700"`
			MusicList []*Song `json:"musicList"`
		} `json:"data"`
	}

	PlayListsResponse struct {
		CommonResponse
		CurTime   int64     `json:"curTime"`
		Data      PlayLists `json:"data"`
		ProfileID string    `json:"profileId"`
		ReqID     string    `json:"reqId"`
	}
	PlayModel struct {
		Img          string `json:"img"`
		Uname        string `json:"uname"`
		LosslessMark string `json:"lossless_mark"`
		Favorcnt     string `json:"favorcnt"`
		Isnew        string `json:"isnew"`
		Extend       string `json:"extend"`
		UID          string `json:"uid"`
		Total        int64  `json:"total,string"`
		Commentcnt   string `json:"commentcnt"`
		Imgscript    string `json:"imgscript"`
		Digest       string `json:"digest"`
		Name         string `json:"name"`
		Listencnt    int64  `json:"listencnt,string,omitempty"`
		ID           string `json:"id"`
		Attribute    string `json:"attribute"`
		RadioID      string `json:"radio_id"`
		Desc         string `json:"desc"`
		Info         string `json:"info"`
	}
	PlayLists struct {
		Total int         `json:"total"`
		Data  []PlayModel `json:"data"`
		Rn    int         `json:"rn"`
		Pn    int         `json:"pn"`
	}

	API struct {
		Client *ghttp.Client
	}
)

func New(client *ghttp.Client) *API {
	return &API{
		Client: client,
	}
}

func Client() *API {
	return std
}

func (c *CommonResponse) errorMessage() string {
	if c.Msg == "" {
		return strconv.Itoa(c.Code)
	}
	return c.Msg
}

func (s *RankResponse) String() string {
	return utils.PrettyJSON(s)
}

func (s *RankListResponse) String() string {
	return utils.PrettyJSON(s)
}

func (s *SearchSongsResponse) String() string {
	return utils.PrettyJSON(s)
}

func (s *SongResponse) String() string {
	return utils.PrettyJSON(s)
}

func (s *SongURLResponse) String() string {
	return utils.PrettyJSON(s)
}

func (s *SongLyricResponse) String() string {
	return utils.PrettyJSON(s)
}

func (a *ArtistInfoResponse) String() string {
	return utils.PrettyJSON(a)
}

func (a *ArtistSongsResponse) String() string {
	return utils.PrettyJSON(a)
}

func (a *AlbumResponse) String() string {
	return utils.PrettyJSON(a)
}

func (p *PlaylistResponse) String() string {
	return utils.PrettyJSON(p)
}

func (a *API) SendRequest(req *ghttp.Request) (*ghttp.Response, error) {
	// csrf 必须跟 kw_token 保持一致
	csrf := "0"
	cookie, err := a.Client.Cookie(req.URL.String(), "kw_token")
	if err != nil {
		req.AddCookie(&http.Cookie{
			Name:  "kw_token",
			Value: csrf,
		})
	} else {
		csrf = cookie.Value
	}

	headers := ghttp.Headers{
		"Origin":  "http://www.kuwo.cn",
		"Referer": "http://www.kuwo.cn",
		"csrf":    csrf,
	}
	req.SetHeaders(headers)
	return a.Client.Do(req)
}

func (a *API) patchSongsURL(ctx context.Context, br int, songs ...*Song) {
	c := concurrency.New(32)
	for _, s := range songs {
		if ctx.Err() != nil {
			break
		}

		c.Add(1)
		go func(s *Song) {
			url, err := a.GetSongURL(ctx, s.RId, br)
			if err == nil {
				s.URL = url
			}
			c.Done()
		}(s)
	}
	c.Wait()
}

func (a *API) patchSongsLyric(ctx context.Context, songs ...*Song) {
	c := concurrency.New(32)
	for _, s := range songs {
		if ctx.Err() != nil {
			break
		}

		c.Add(1)
		go func(s *Song) {
			lyric, err := a.GetSongLyric(ctx, s.RId)
			if err == nil {
				s.Lyric = lyric
			}
			c.Done()
		}(s)
	}
	c.Wait()
}

func translate(src ...*Song) []*api.Song {
	songs := make([]*api.Song, len(src))
	for i, s := range src {
		songs[i] = &api.Song{
			Id:        strconv.Itoa(s.RId),
			Name:      strings.TrimSpace(s.Name),
			Artist:    strings.TrimSpace(strings.ReplaceAll(s.Artist, "&", "/")),
			Album:     strings.TrimSpace(s.Album),
			PicURL:    s.AlbumPic,
			Lyric:     s.Lyric,
			ListenURL: s.URL,
		}
	}
	return songs
}
