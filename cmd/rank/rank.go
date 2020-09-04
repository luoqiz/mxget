package rank

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/winterssy/glog"
	"github.com/winterssy/mxget/internal/settings"
	"github.com/winterssy/mxget/pkg/provider"
)

var (
	from string
)

var CmdRank = &cobra.Command{
	Use:   "rank",
	Short: "Get rank from the specified music platform",
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

	fmt.Printf("Search rank from [%s]...\n\n", provider.GetDesc(platform))
	result, err := client.GetRank(context.Background())
	if err != nil {
		glog.Fatal(err)
	}

	var sb strings.Builder
	for i, s := range result {
		fmt.Fprintf(&sb, "[%02d] %s - %s - %s - %s\n", i+1, s.Sourceid, s.Name, s.Intro, s.ID)
	}
	fmt.Println(sb.String())

	if from != "" {
		fmt.Printf("Command: mxget rank --from %s \n", from)
	}
}

func init() {
	CmdRank.Flags().StringVar(&from, "from", "", "music platform")
	CmdRank.Run = Run
}
