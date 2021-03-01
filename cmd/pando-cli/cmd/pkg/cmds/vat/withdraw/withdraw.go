package withdraw

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/internal/cfg"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/pkg/cmds/actions"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/pkg/cmds/pay"
	"github.com/fox-one/pando/core"
	"github.com/fox-one/pando/pkg/mtg/types"
	"github.com/fox-one/pando/pkg/number"
	"github.com/fox-one/pando/pkg/uuid"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "withdraw",
		Args: cobra.ExactValidArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			token := cfg.GetAuthToken()
			user, err := mixin.UserMe(ctx, token)
			if err != nil {
				return err
			}
			// follow id
			follow := uuid.New()

			vatID := args[0]
			dink := number.Decimal(args[1])
			memo, err := actions.Tx(
				core.ActionVatWithdraw,
				types.UUID(user.UserID),
				types.UUID(follow),
				types.UUID(vatID),
				dink,
			)
			if err != nil {
				return err
			}

			cmd.Println("tx follow id:", follow)
			return pay.Request(ctx, pay.CNB, number.One, memo)
		},
	}

	return cmd
}