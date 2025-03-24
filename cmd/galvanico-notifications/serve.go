package galvanico_notification

import (
	"fmt"
	"github.com/spf13/cobra"
)

// NotificationCmd represents the serve command
var NotificationCmd = &cobra.Command{
	Use:   "notification",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	NotificationCmd.AddCommand(sendCmd)
	NotificationCmd.AddCommand(runCmd)
}
