package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Messages and activity comments",
}

var chatListCmd = &cobra.Command{
	Use:   "list",
	Short: "List chats",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/chats"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var chatGetCmd = &cobra.Command{
	Use:   "get <chat-id>",
	Short: "Get a chat",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/chats/%s", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var chatMessagesCmd = &cobra.Command{
	Use:   "messages <chat-id>",
	Short: "List messages in a chat",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		limit, _ := cmd.Flags().GetInt("limit")
		params := url.Values{}
		if limit > 0 {
			params.Set("limit", fmt.Sprintf("%d", limit))
		}
		path := fmt.Sprintf("/chats/%s/messages", args[0])
		if len(params) > 0 {
			path += "?" + params.Encode()
		}
		data, err := client.Get(path)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var chatSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message",
	Long:  "Send a message. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw("/chats/send-message", body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var chatDeleteMessageCmd = &cobra.Command{
	Use:   "delete-message <chat-id> <message-id>",
	Short: "Delete a message",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(fmt.Sprintf("/chats/%s/messages/%s", args[0], args[1]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var chatActivityCommentsCmd = &cobra.Command{
	Use:   "activity-comments <activity-id>",
	Short: "List comments on an activity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/activity/%s/messages", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var chatCommentCmd = &cobra.Command{
	Use:   "comment <activity-id>",
	Short: "Add a comment to an activity",
	Long:  "Add activity comment. Pass JSON body via stdin.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(fmt.Sprintf("/activity/%s/messages", args[0]), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	chatMessagesCmd.Flags().Int("limit", 30, "max messages to return")

	chatCmd.AddCommand(chatListCmd)
	chatCmd.AddCommand(chatGetCmd)
	chatCmd.AddCommand(chatMessagesCmd)
	chatCmd.AddCommand(chatSendCmd)
	chatCmd.AddCommand(chatDeleteMessageCmd)
	chatCmd.AddCommand(chatActivityCommentsCmd)
	chatCmd.AddCommand(chatCommentCmd)
	rootCmd.AddCommand(chatCmd)
}
