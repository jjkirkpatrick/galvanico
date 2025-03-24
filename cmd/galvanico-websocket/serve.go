package galvanico_websocket

import (
	"fmt"
	"github.com/spf13/cobra"
)

// WsCmd represents the serve command
var WsCmd = &cobra.Command{
	Use:   "ws",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	WsCmd.AddCommand(sendCmd)
	WsCmd.AddCommand(runCmd)
}
