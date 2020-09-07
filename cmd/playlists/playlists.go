package playlist

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winterssy/glog"
	"github.com/winterssy/mxget/internal/settings"
	"github.com/winterssy/mxget/pkg/provider"
	"github.com/winterssy/mxget/pkg/utils"
	"strconv"
)

var (
	playlistId string
	from       string
	page       int
	pageSize   int
)

var CmdPlaylists = &cobra.Command{
	Use:   "playlists",
	Short: "Fetch play list",
}

func Run(cmd *cobra.Command, args []string) {

	platform := settings.Cfg.Platform
	if from != "" {
		platform = from
	}

	client, err := provider.GetClient(platform)
	if err != nil {
		glog.Fatal(err)
	}
	if page == 0 {
		pageStr := utils.Input("page")
		page, _ = strconv.Atoi(pageStr)
		fmt.Println()
	}

	if pageSize == 0 {
		pageSizeStr := utils.Input("pageSize")
		pageSize, _ = strconv.Atoi(pageSizeStr)
		fmt.Println()
	}
	glog.Infof("Fetch playlist from [%s]", provider.GetDesc(platform))
	ctx := context.Background()
	playlist, err := client.GetPlayLists(ctx, page, pageSize)
	if err != nil {
		glog.Fatal(err)
	}

	for i, p := range playlist {
		glog.Infof("%v - %s - %s", i, p.ID, p.Name)
	}
}

func init() {
	CmdPlaylists.Flags().StringVar(&from, "from", "", "music platform")
	CmdPlaylists.Flags().IntVarP(&page, "page", "p", 0, "page")
	CmdPlaylists.Flags().IntVarP(&pageSize, "pageSize", "r", 0, "pageSize")
	CmdPlaylists.Run = Run
}
