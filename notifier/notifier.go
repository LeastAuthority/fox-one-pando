package notifier

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fox-one/mixin-sdk-go"
	"github.com/fox-one/pando/core"
	"github.com/fox-one/pando/pkg/uuid"
	"github.com/fox-one/pando/service/asset"
	"github.com/fox-one/pkg/logger"
	"github.com/fox-one/pkg/text/localizer"
)

func New(
	system *core.System,
	assetz core.AssetService,
	messages core.MessageStore,
	vats core.VaultStore,
	cats core.CollateralStore,
	users core.UserStore,
	i18n *localizer.Localizer,
) core.Notifier {
	return &notifier{
		system:   system,
		assetz:   asset.Cache(assetz),
		messages: messages,
		vats:     vats,
		cats:     cats,
		users:    users,
		i18n:     i18n,
	}
}

type notifier struct {
	system   *core.System
	assetz   core.AssetService
	messages core.MessageStore
	vats     core.VaultStore
	cats     core.CollateralStore
	users    core.UserStore
	i18n     *localizer.Localizer
}

func (n *notifier) localize(id, lang string, args ...interface{}) string {
	l := n.i18n
	if lang != "" {
		l = localizer.WithLanguage(l, lang)
	}

	s := l.LocalizeOr(id, id, args...)
	return strings.TrimSpace(s)
}

func (n *notifier) Auth(ctx context.Context, user *core.User) error {
	msg := n.localize("login_done", user.Lang)
	req := &mixin.MessageRequest{
		ConversationID: mixin.UniqueConversationID(n.system.ClientID, user.MixinID),
		RecipientID:    user.MixinID,
		MessageID:      uuid.Modify(user.MixinID, user.AccessToken),
		Category:       mixin.MessageCategoryPlainText,
		Data:           base64.StdEncoding.EncodeToString([]byte(msg)),
	}

	return n.messages.Create(ctx, []*core.Message{core.BuildMessage(req)})
}

func (n *notifier) Transaction(ctx context.Context, tx *core.Transaction) error {
	if tx.UserID == "" {
		return nil
	}

	user, err := n.users.Find(ctx, tx.UserID)
	if err != nil {
		return err
	}

	action := n.localize("Action"+tx.Action.String(), user.Lang)
	data := TxData{
		Action:  action,
		Message: tx.Message,
	}

	id := "tx_abort"
	if tx.Status == core.TransactionStatusOk {
		switch tx.Action {
		case core.ActionVatOpen, core.ActionVatDeposit, core.ActionVatWithdraw, core.ActionVatPayback, core.ActionVatGenerate:
			if err := n.handleVatTx(ctx, tx, user, &data); err != nil {
				return err
			}
		}

		id = "tx_ok"
	}

	msg := n.localize(id, user.Lang, data)
	req := &mixin.MessageRequest{
		ConversationID: mixin.UniqueConversationID(n.system.ClientID, tx.UserID),
		RecipientID:    tx.UserID,
		MessageID:      uuid.Modify(tx.TraceID, "notify"),
		Category:       mixin.MessageCategoryPlainText,
		Data:           base64.StdEncoding.EncodeToString([]byte(msg)),
	}

	return n.messages.Create(ctx, []*core.Message{core.BuildMessage(req)})
}

func (n *notifier) Snapshot(ctx context.Context, transfer *core.Transfer, signedTx string) error {
	log := logger.FromContext(ctx)

	tx, err := mixin.TransactionFromRaw(signedTx)
	if err != nil {
		log.WithError(err).Debugln("decode transaction from raw tx failed")
		return nil
	}

	hash, err := tx.TransactionHash()
	if err != nil {
		return nil
	}

	traceID := mixinRawTransactionTraceId(hash.String(), 0)

	if len(transfer.Opponents) != 1 {
		log.Debugln("transfer opponents is not 1")
		return nil
	}

	coin, err := n.assetz.Find(ctx, transfer.AssetID)
	if err != nil {
		return err
	}

	card := mixin.AppCardMessage{
		AppID:       n.system.ClientID,
		IconURL:     coin.Logo,
		Title:       transfer.Amount.String(),
		Description: coin.Symbol,
		Action:      mixin.URL.Snapshots("", traceID),
	}
	data, _ := json.Marshal(card)

	recipientID := transfer.Opponents[0]
	req := &mixin.MessageRequest{
		ConversationID: mixin.UniqueConversationID(n.system.ClientID, recipientID),
		RecipientID:    recipientID,
		MessageID:      traceID,
		Category:       mixin.MessageCategoryAppCard,
		Data:           base64.StdEncoding.EncodeToString(data),
	}

	return n.messages.Create(ctx, []*core.Message{core.BuildMessage(req)})
}

func (n *notifier) VaultUnsafe(ctx context.Context, cat *core.Collateral, vault *core.Vault) error {
	user, err := n.users.Find(ctx, vault.UserID)
	if err != nil {
		return err
	}

	msg := n.localize("vat_unsafe_warn", user.Lang, map[string]interface{}{
		"ID":   vault.ID,
		"Name": cat.Name,
	})

	req := &mixin.MessageRequest{
		ConversationID: mixin.UniqueConversationID(n.system.ClientID, user.MixinID),
		RecipientID:    user.MixinID,
		MessageID:      uuid.Modify(vault.TraceID, fmt.Sprintf("versions_%d_%d", cat.Version, vault.Version)),
		Category:       mixin.MessageCategoryPlainText,
		Data:           base64.StdEncoding.EncodeToString([]byte(msg)),
	}

	return n.messages.Create(ctx, []*core.Message{core.BuildMessage(req)})
}
