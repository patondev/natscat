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

var publishSubject string
var publishMessage string

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "To publish a message with specific subject",
	Long: `To publish a message with specific subject`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Conneting to %v\n",natsAddress)
		npub := nats.NatsClass{DefaultURL: natsAddress, PubSubject: publishSubject, Message: publishMessage}
		npub.Publish()
	},
}

func init() {
	rootCmd.AddCommand(publishCmd)
	rf := publishCmd.Flags()
	rf.StringVarP(&publishSubject, "subject", "s", "", "Publish subject (required)")
	rf.StringVarP(&publishMessage, "message", "m", "", "Publish message (required)")
	cobra.MarkFlagRequired(rf,"subject")
	cobra.MarkFlagRequired(rf,"message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
