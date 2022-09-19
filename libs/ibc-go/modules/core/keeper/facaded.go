package keeper

import (
	"context"
	"errors"

	types2 "github.com/okex/exchain/libs/tendermint/types"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	clienttypes "github.com/okex/exchain/libs/ibc-go/modules/core/02-client/types"
	connectiontypes "github.com/okex/exchain/libs/ibc-go/modules/core/03-connection/types"
	channeltyeps "github.com/okex/exchain/libs/ibc-go/modules/core/04-channel/types"
)

var _ IBCServerKeeper = (*FacadedKeeper)(nil)

var errMisSpecificKeeper = errors.New("mis ")

var (
	V4Register KeepRegister = func() (KeeperSelector, int64) {
		return func(ctx sdk.Context) (int64, bool) {
			if types2.HigherThanVenus3(ctx.BlockHeight()) {
				return types2.GetVenus3Height(), true
			}
			return 0, false
		}, types2.GetVenus3Height()
	}
)

type KeepRegister func() (KeeperSelector, int64)
type KeeperSelector func(ctx sdk.Context) (int64, bool)

type IBCServerKeeper interface {
	channeltyeps.QueryServer
	channeltyeps.MsgServer
	clienttypes.MsgServer
	connectiontypes.MsgServer

	GetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) (string, bool)
	GetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) []byte
}

type FacadedKeeper struct {
	V2Keeper  *Keeper
	keepers   map[int64]IBCServerKeeper
	selectors []KeeperSelector
}

func NewFacadedKeeper(v2Keeper *Keeper) *FacadedKeeper {
	ret := &FacadedKeeper{
		keepers: make(map[int64]IBCServerKeeper),
	}
	ret.V2Keeper = v2Keeper

	return ret
}

func (f *FacadedKeeper) RegisterKeeper(k IBCServerKeeper, reg KeepRegister) {
	selector, h := reg()
	if _, exist := f.keepers[h]; exist {
		// TODO,错误码要不要改下?
		panic("program error")
	}
	f.keepers[h] = k
	f.selectors = append(f.selectors, selector)
}

func (f *FacadedKeeper) GetPacketCommitment(ctx sdk.Context, portID, channelID string, sequence uint64) []byte {
	k := f.doGetByCtx(ctx)
	return k.GetPacketCommitment(ctx, portID, channelID, sequence)
}

func (f *FacadedKeeper) GetPacketReceipt(ctx sdk.Context, portID, channelID string, sequence uint64) (string, bool) {
	k := f.doGetByCtx(ctx)
	return k.GetPacketReceipt(ctx, portID, channelID, sequence)
}

func (f *FacadedKeeper) Channel(goCtx context.Context, request *channeltyeps.QueryChannelRequest) (*channeltyeps.QueryChannelResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.Channel(goCtx, request)
}

func (f *FacadedKeeper) Channels(goCtx context.Context, request *channeltyeps.QueryChannelsRequest) (*channeltyeps.QueryChannelsResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.Channels(goCtx, request)
}

func (f *FacadedKeeper) ConnectionChannels(goCtx context.Context, request *channeltyeps.QueryConnectionChannelsRequest) (*channeltyeps.QueryConnectionChannelsResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ConnectionChannels(goCtx, request)
}

func (f *FacadedKeeper) ChannelClientState(goCtx context.Context, request *channeltyeps.QueryChannelClientStateRequest) (*channeltyeps.QueryChannelClientStateResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelClientState(goCtx, request)
}

func (f *FacadedKeeper) ChannelConsensusState(goCtx context.Context, request *channeltyeps.QueryChannelConsensusStateRequest) (*channeltyeps.QueryChannelConsensusStateResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelConsensusState(goCtx, request)
}

func (f *FacadedKeeper) PacketCommitment(goCtx context.Context, request *channeltyeps.QueryPacketCommitmentRequest) (*channeltyeps.QueryPacketCommitmentResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.PacketCommitment(goCtx, request)
}

func (f *FacadedKeeper) PacketCommitments(goCtx context.Context, request *channeltyeps.QueryPacketCommitmentsRequest) (*channeltyeps.QueryPacketCommitmentsResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.PacketCommitments(goCtx, request)
}

func (f *FacadedKeeper) PacketReceipt(goCtx context.Context, request *channeltyeps.QueryPacketReceiptRequest) (*channeltyeps.QueryPacketReceiptResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.PacketReceipt(goCtx, request)
}

func (f *FacadedKeeper) PacketAcknowledgement(goCtx context.Context, request *channeltyeps.QueryPacketAcknowledgementRequest) (*channeltyeps.QueryPacketAcknowledgementResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.PacketAcknowledgement(goCtx, request)
}

func (f *FacadedKeeper) PacketAcknowledgements(goCtx context.Context, request *channeltyeps.QueryPacketAcknowledgementsRequest) (*channeltyeps.QueryPacketAcknowledgementsResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.PacketAcknowledgements(goCtx, request)
}

func (f *FacadedKeeper) UnreceivedPackets(goCtx context.Context, request *channeltyeps.QueryUnreceivedPacketsRequest) (*channeltyeps.QueryUnreceivedPacketsResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.UnreceivedPackets(goCtx, request)
}

func (f *FacadedKeeper) UnreceivedAcks(goCtx context.Context, request *channeltyeps.QueryUnreceivedAcksRequest) (*channeltyeps.QueryUnreceivedAcksResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.UnreceivedAcks(goCtx, request)
}

func (f *FacadedKeeper) NextSequenceReceive(goCtx context.Context, request *channeltyeps.QueryNextSequenceReceiveRequest) (*channeltyeps.QueryNextSequenceReceiveResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.NextSequenceReceive(goCtx, request)
}

func (f *FacadedKeeper) ChannelOpenInit(goCtx context.Context, init *channeltyeps.MsgChannelOpenInit) (*channeltyeps.MsgChannelOpenInitResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelOpenInit(goCtx, init)
}

func (f *FacadedKeeper) ChannelOpenTry(goCtx context.Context, try *channeltyeps.MsgChannelOpenTry) (*channeltyeps.MsgChannelOpenTryResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelOpenTry(goCtx, try)
}

func (f *FacadedKeeper) ChannelOpenAck(goCtx context.Context, ack *channeltyeps.MsgChannelOpenAck) (*channeltyeps.MsgChannelOpenAckResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelOpenAck(goCtx, ack)
}

func (f *FacadedKeeper) ChannelOpenConfirm(goCtx context.Context, confirm *channeltyeps.MsgChannelOpenConfirm) (*channeltyeps.MsgChannelOpenConfirmResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelOpenConfirm(goCtx, confirm)
}

func (f *FacadedKeeper) ChannelCloseInit(goCtx context.Context, init *channeltyeps.MsgChannelCloseInit) (*channeltyeps.MsgChannelCloseInitResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelCloseInit(goCtx, init)
}

func (f *FacadedKeeper) ChannelCloseConfirm(goCtx context.Context, confirm *channeltyeps.MsgChannelCloseConfirm) (*channeltyeps.MsgChannelCloseConfirmResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ChannelCloseConfirm(goCtx, confirm)
}

func (f *FacadedKeeper) RecvPacket(goCtx context.Context, packet *channeltyeps.MsgRecvPacket) (*channeltyeps.MsgRecvPacketResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.RecvPacket(goCtx, packet)
}

func (f *FacadedKeeper) Timeout(goCtx context.Context, timeout *channeltyeps.MsgTimeout) (*channeltyeps.MsgTimeoutResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.Timeout(goCtx, timeout)
}

func (f *FacadedKeeper) TimeoutOnClose(goCtx context.Context, onClose *channeltyeps.MsgTimeoutOnClose) (*channeltyeps.MsgTimeoutOnCloseResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.TimeoutOnClose(goCtx, onClose)
}

func (f *FacadedKeeper) Acknowledgement(goCtx context.Context, acknowledgement *channeltyeps.MsgAcknowledgement) (*channeltyeps.MsgAcknowledgementResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.Acknowledgement(goCtx, acknowledgement)
}

func (f *FacadedKeeper) CreateClient(goCtx context.Context, client *clienttypes.MsgCreateClient) (*clienttypes.MsgCreateClientResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.CreateClient(goCtx, client)
}

func (f *FacadedKeeper) UpdateClient(goCtx context.Context, client *clienttypes.MsgUpdateClient) (*clienttypes.MsgUpdateClientResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.UpdateClient(goCtx, client)
}

func (f *FacadedKeeper) UpgradeClient(goCtx context.Context, client *clienttypes.MsgUpgradeClient) (*clienttypes.MsgUpgradeClientResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.UpgradeClient(goCtx, client)
}

func (f *FacadedKeeper) SubmitMisbehaviour(goCtx context.Context, misbehaviour *clienttypes.MsgSubmitMisbehaviour) (*clienttypes.MsgSubmitMisbehaviourResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.SubmitMisbehaviour(goCtx, misbehaviour)
}

func (f *FacadedKeeper) ConnectionOpenInit(goCtx context.Context, init *connectiontypes.MsgConnectionOpenInit) (*connectiontypes.MsgConnectionOpenInitResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ConnectionOpenInit(goCtx, init)
}

func (f *FacadedKeeper) ConnectionOpenTry(goCtx context.Context, try *connectiontypes.MsgConnectionOpenTry) (*connectiontypes.MsgConnectionOpenTryResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ConnectionOpenTry(goCtx, try)
}

func (f *FacadedKeeper) ConnectionOpenAck(goCtx context.Context, ack *connectiontypes.MsgConnectionOpenAck) (*connectiontypes.MsgConnectionOpenAckResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ConnectionOpenAck(goCtx, ack)
}

func (f *FacadedKeeper) ConnectionOpenConfirm(goCtx context.Context, confirm *connectiontypes.MsgConnectionOpenConfirm) (*connectiontypes.MsgConnectionOpenConfirmResponse, error) {
	k := f.getHeightKeeper(goCtx)
	return k.ConnectionOpenConfirm(goCtx, confirm)
}

func (f *FacadedKeeper) getHeightKeeper(goCtx context.Context) IBCServerKeeper {
	ctx := sdk.UnwrapSDKContext(goCtx)
	return f.doGetByCtx(ctx)
}

func (f *FacadedKeeper) doGetByCtx(ctx sdk.Context) IBCServerKeeper {
	for _, selector := range f.selectors {
		hh, ok := selector(ctx)
		if ok {
			return f.doGet(hh)
		}
	}
	return f.V2Keeper
}

func (f *FacadedKeeper) doGet(h int64) IBCServerKeeper {
	ret, exist := f.keepers[h]
	if !exist {
		return f.V2Keeper
	}
	return ret
}

func (f *FacadedKeeper) GetIbcEnabled(ctx sdk.Context) bool {
	return f.V2Keeper.GetIbcEnabled(ctx)
}
