package flip

import (
	"github.com/fox-one/pando/core"
	"github.com/fox-one/pando/pkg/maker"
	"github.com/shopspring/decimal"
)

func Tend(r *maker.Request, c *core.Collateral, f *core.Flip, lot decimal.Decimal) error {
	if err := require(f.TicFinished(r.Now) == false, "finished-tic"); err != nil {
		return err
	}

	if err := require(f.EndFinished(r.Now), "finished-end"); err != nil {
		return err
	}

	assetID, bid := r.AssetID, r.Amount
	if err := require(assetID == c.Dai && bid.LessThanOrEqual(f.Tab), "bid-not-match"); err != nil {
		return err
	}

	if err := require(f.Lot.Equal(lot), "lot-not-match"); err != nil {
		return err
	}

	if err := require(bid.GreaterThan(f.Bid), "bid-not-higher"); err != nil {
		return err
	}

	if err := require(bid.Equal(f.Tab) || bid.GreaterThanOrEqual(f.Bid.Mul(c.Beg)), "insufficient-increase"); err != nil {
		return err
	}

	return nil
}
