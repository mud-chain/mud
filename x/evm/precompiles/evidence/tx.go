package evidence

import (
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/evmos/v12/x/evm/types"
)

const (
	SubmitEvidenceGas = 60_000

	SubmitEvidenceMethodName = "submitEvidence"

	SubmitEvidenceEventName = "SubmitEvidence"
)

func (c *Contract) SubmitEvidence(ctx sdk.Context, evm *vm.EVM, contract *vm.Contract, readonly bool) ([]byte, error) {
	if readonly {
		return nil, types.ErrReadOnly
	}

	if evm.Origin != contract.Caller() {
		return nil, types.ErrInvalidCaller
	}

	method := MustMethod(SubmitEvidenceMethodName)

	// parse args
	var args SubmitEvidenceArgs
	err := types.ParseMethodArgs(method, &args, contract.Input[4:])
	if err != nil {
		return nil, err
	}

	evidence := &evidencetypes.Equivocation{
		Height:           args.Height,
		Power:            args.Power,
		Time:             time.UnixMilli(args.Time),
		ConsensusAddress: args.GetConsAddress().String(),
	}

	msg, err := evidencetypes.NewMsgSubmitEvidence(contract.Caller().Bytes(), evidence)
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	server := evidencekeeper.NewMsgServerImpl(c.evidenceKeeper)
	res, err := server.SubmitEvidence(ctx, msg)
	if err != nil {
		return nil, err
	}

	hash := hex.EncodeToString(res.Hash)

	// add submit evidence log
	if err := c.AddLog(
		evm,
		MustEvent(SubmitEvidenceEventName),
		[]common.Hash{common.BytesToHash(contract.Caller().Bytes())},
		hash,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(hash)
}
