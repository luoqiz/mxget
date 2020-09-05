package ranklist

import (
	"context"
	"fmt"
	"github.com/winterssy/mxget/pkg/utils"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/winterssy/glog"
	"github.com/winterssy/mxget/internal/settings"
	"github.com/winterssy/mxget/pkg/provider"
)

var (
	bangId   string
	from     string
	page     int
	pageSize int
)

var CmdRankList = &cobra.Command{
	Use:   "ranklist",
	Short: "Get ranklist from the specified music platform",
}

func Run(cmd *cobra.Command, args []string) {

	if bangId == "" {
		bangId = utils.Input("bangId")
		fmt.Println()
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

	platform := settings.Cfg.Platform
	if from != "" {
		platform = from
	}

	client, err := provider.GetClient(platform)
	if err != nil {
		glog.Fatal(err)
	}

	fmt.Printf("Search rank list from [%s]...\n\n", provider.GetDesc(platform))
	result, err := client.GetRankList(context.Background(), bangId, page, pageSize)
	if err != nil {
		glog.Fatal(err)
	}

	var sb strings.Builder
	for i, s := range result {
		fmt.Fprintf(&sb, "[%02d] %s - %s - %s - %s\n", i+1, s.Id, s.Name, s.Album, s.ListenURL)
	}
	fmt.Println(sb.String())

	if from != "" {
		fmt.Printf("Command: mxget rank --from %s \n", from)
	}
}

func init() {
	CmdRankList.Flags().StringVarP(&bangId, "bangId", "b", "", "rank id")
	CmdRankList.Flags().StringVar(&from, "from", "", "music platform")
	CmdRankList.Flags().IntVarP(&page, "page", "p", 0, "page")
	CmdRankList.Flags().IntVarP(&pageSize, "pageSize", "r", 0, "pageSize")
	CmdRankList.Run = Run
}
