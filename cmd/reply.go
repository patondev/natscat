/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/patondev/natscat/internal/nats"
)

var replySubject string
var replyMessage string
var queueGroup string

// replyCmd represents the reply command
var replyCmd = &cobra.Command{
	Use:   "rep",
	Short: "To reply to the request's subject with specific message",
	Long: `To reply to the request's subject with specific message`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Conneting to %v\n",natsAddress)
		nreply := nats.NatsClass{DefaultURL: natsAddress, ReplySubject: replySubject, ReplyMessage: replyMessage, QueueGroupName: queueGroup}
		nreply.Reply()
	},
}

func init() {
	rootCmd.AddCommand(replyCmd)
	rf := replyCmd.Flags()
	rf.StringVarP(&replySubject, "subject", "s", "", "Reply subject (required)")
	rf.StringVarP(&replyMessage, "message", "m", "", "Reply message (required)")
	rf.StringVar(&queueGroup, "qgroup", "default", "Queue Group")
	cobra.MarkFlagRequired(rf,"subject")
	cobra.MarkFlagRequired(rf,"message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
