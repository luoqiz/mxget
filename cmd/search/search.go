package search

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/winterssy/glog"
	"github.com/winterssy/mxget/internal/settings"
	"github.com/winterssy/mxget/pkg/provider"
	"github.com/winterssy/mxget/pkg/utils"
)

var (
	keyword  string
	from     string
	page     int
	pageSize int
)

var CmdSearch = &cobra.Command{
	Use:   "search",
	Short: "Search songs from the specified music platform",
}

func Run(cmd *cobra.Command, args []string) {
	if keyword == "" {
		keyword = utils.Input("Keyword")
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

	fmt.Printf("Search %q from [%s]...\n\n", keyword, provider.GetDesc(platform))
	result, err := client.SearchSongs(context.Background(), keyword, page, pageSize)
	if err != nil {
		glog.Fatal(err)
	}

	var sb strings.Builder
	for i, s := range result {
		fmt.Fprintf(&sb, "[%02d] %s - %s - %s\n", i+1, s.Name, s.Artist, s.Id)
	}
	fmt.Println(sb.String())

	if from != "" {
		fmt.Printf("Command: mxget song --from %s --id <song id>\n", from)
	} else {
		fmt.Println("Command: mxget song --id <song id>")
	}
}

func init() {
	CmdSearch.Flags().StringVarP(&keyword, "keyword", "k", "", "search keyword")
	CmdSearch.Flags().StringVar(&from, "from", "", "music platform")
	CmdSearch.Flags().IntVarP(&page, "page", "p", 0, "page")
	CmdSearch.Flags().IntVarP(&pageSize, "pageSize", "r", 0, "pageSize")
	CmdSearch.Run = Run
}
