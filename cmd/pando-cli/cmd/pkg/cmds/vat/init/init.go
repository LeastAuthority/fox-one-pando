package init

import (
	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/internal/call"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/internal/cfg"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/pkg/cmds/actions"
	"github.com/fox-one/pando/cmd/pando-cli/cmd/pkg/cmds/pay"
	"github.com/fox-one/pando/core"
	"github.com/fox-one/pando/handler/rpc/api"
	"github.com/fox-one/pando/pkg/mtg/types"
	"github.com/fox-one/pando/pkg/number"
	"github.com/fox-one/pando/pkg/uuid"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "init <collateral id> <deposit> <generate>",
		Args: cobra.ExactValidArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			token := cfg.GetAuthToken()
			user, err := mixin.UserMe(ctx, token)
			if err != nil {
				return err
			}
			// follow id
			follow := uuid.New()

			catID := args[0]
			cat, err := call.RPC().FindCollateral(ctx, &api.Req_FindCollateral{Id: catID})
			if err != nil {
				return err
			}

			dink := number.Decimal(args[1])
			debt := number.Decimal(args[2])
			memo, err := actions.Tx(
				core.ActionVatInit,
				types.UUID(user.UserID),
				types.UUID(follow),
				types.UUID(catID),
				debt,
			)
			if err != nil {
				return err
			}

			cmd.Println("tx follow id:", follow)
			return pay.Request(ctx, cat.Gem, dink, memo)
		},
	}

	return cmd
}