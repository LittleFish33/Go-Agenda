package cmd

import (
	"Go-Agenda/entity"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// deleteAllMeetingCmd represents the deleteAllMeeting command
var deleteAllMeetingCmd = &cobra.Command{
	Use:   "deleteAllMeeting",
	Short: "Delete all meeting sponsor by you",
	Long:  "Delete all meeting sponsor by you，only the meeting sponsored by you can be deleted!",
	Run: func(cmd *cobra.Command, args []string) {
		//判断是否已经登录
		if entity.IsLogin() {
			var remainMeeting []entity.Meeting
			for _, meet := range entity.GetMeetings() {
				if meet.Sponsor != entity.Loginuser.Name {
					remainMeeting = append(remainMeeting, meet)
				}
			}
			if len(entity.GetMeetings()) == len(remainMeeting) {
				fmt.Println("You have no meeting to delete!")
				log.Println("You have no meeting to delete!")
				return
			}
			//写入文件
			entity.Meetings = remainMeeting
			entity.WriteJson()
			fmt.Println("Delete All Meeting Successfully!")
			log.Println("Delete All Meeting Successfully!")
		} else {
			fmt.Println("You must login first!")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteAllMeetingCmd)
}
