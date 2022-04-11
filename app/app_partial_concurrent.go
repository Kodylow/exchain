package app

import (
	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/libs/cosmos-sdk/x/auth"
	authante "github.com/okex/exchain/libs/cosmos-sdk/x/auth/ante"
	evmtypes "github.com/okex/exchain/x/evm/types"
)

// getTxFeeAndFromHandler get tx fee and from
func getTxFeeAndFromHandler(ak auth.AccountKeeper) sdk.GetTxFeeAndFromHandler {
	return func(ctx sdk.Context, tx sdk.Tx) (fee sdk.Coins, isEvm bool, from string, to string) {
		if evmTx, ok := tx.(*evmtypes.MsgEthereumTx); ok {
			isEvm = true
			_ = evmTx.VerifySig(evmTx.ChainID(), ctx.BlockHeight())
			fee = evmTx.GetFee()
			feePayer := evmTx.FeePayer(ctx)//.AccountAddress()
			//feeReceiver := evmTx.To()
			//if feeReceiver != nil {
			//	to = string(feeReceiver.Bytes())
			//}
			feePayerAcc := ak.GetAccount(ctx, feePayer)
			from = feePayerAcc.GetAddress().String()
			//from = string(feePayerAcc.GetAddress().Bytes())//.String()//hex.EncodeToString(feePayerAcc.GetAddress())
		} else if feeTx, ok := tx.(authante.FeeTx); ok {
			fee = feeTx.GetFee()
			feePayer := feeTx.FeePayer(ctx)
			//from = ak.GetAccount(ctx, feePayer)
			feePayerAcc := ak.GetAccount(ctx, feePayer)
			from = feePayerAcc.GetAddress().String()
			//from = string(feePayerAcc.GetAddress().Bytes())//.String()// ex17xpfvakm2amg962yls6f84z3kell8c5lcs49z2
			//hex.EncodeToString(feePayerAcc.GetAddress())// f1829676db577682e944fc3493d451b67ff3e29f
		}

		return
	}
}
