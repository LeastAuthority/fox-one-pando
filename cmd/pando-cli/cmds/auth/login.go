package auth

import (
	"fmt"
	"net/url"

	"github.com/fox-one/pando/cmd/pando-cli/internal/call"
	"github.com/fox-one/pando/cmd/pando-cli/internal/cfg"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func NewLoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "login",
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				clientID := cfg.GetAuthClient()
				if clientID == "" {
					return fmt.Errorf("oauth client id not set, run pd use {host} first")
				}

				return requestMixinOauth(clientID)
			}

			r, err := call.R(cmd.Context()).SetBody(map[string]interface{}{
				"code": args[0],
			}).Post("/api/login")
			if err != nil {
				return err
			}

			var body struct {
				ID    string `json:"id,omitempty"`
				Name  string `json:"name,omitempty"`
				Token string `json:"token,omitempty"`
			}

			if err := call.UnmarshalResponse(r, &body); err != nil {
				return err
			}

			cmd.Printf("%s welcome!\n", body.Name)
			cfg.SetAuthToken(body.Token)
			cfg.SetAuthMixinID(body.ID)

			return cfg.Save()
		},
	}

	return cmd
}

func requestMixinOauth(clientID string) error {
	q := url.Values{}
	q.Set("client_id", clientID)
	q.Set("scope", "PROFILE:READ ASSETS:READ SNAPSHOTS:READ")
	q.Set("response_type", "code")
	q.Set("redirect_url", "https://mixin-oauth.firesbox.com/general-callback")

	u := &url.URL{
		Scheme:   "https",
		Host:     "mixin-oauth.firesbox.com",
		RawQuery: q.Encode(),
	}

	return browser.OpenURL(u.String())
}
