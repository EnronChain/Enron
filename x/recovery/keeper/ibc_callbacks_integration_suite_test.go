package keeper_test

import (
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	ibcgotesting "github.com/cosmos/ibc-go/v3/testing"

	ibctesting "github.com/enronchain/enron/v3/ibc/testing"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/enronchain/enron/v3/app"
	claimtypes "github.com/enronchain/enron/v3/x/claims/types"
	inflationtypes "github.com/enronchain/enron/v3/x/inflation/types"
	"github.com/enronchain/enron/v3/x/recovery/types"
)

type IBCTestingSuite struct {
	suite.Suite
	coordinator *ibcgotesting.Coordinator

	// testing chains used for convenience and readability
	EnronChain      *ibcgotesting.TestChain
	IBCOsmosisChain *ibcgotesting.TestChain
	IBCCosmosChain  *ibcgotesting.TestChain

	pathOsmosisEnron  *ibcgotesting.Path
	pathCosmosEnron   *ibcgotesting.Path
	pathOsmosisCosmos *ibcgotesting.Path
}

var s *IBCTestingSuite

func TestIBCTestingSuite(t *testing.T) {
	s = new(IBCTestingSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *IBCTestingSuite) SetupTest() {
	// initializes 3 test chains
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 1, 2)
	suite.EnronChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(1))
	suite.IBCOsmosisChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(2))
	suite.IBCCosmosChain = suite.coordinator.GetChain(ibcgotesting.GetChainID(3))
	suite.coordinator.CommitNBlocks(suite.EnronChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCOsmosisChain, 2)
	suite.coordinator.CommitNBlocks(suite.IBCCosmosChain, 2)

	// Mint coins locked on the enron account generated with secp.
	coinEnron := sdk.NewCoin("aenron", sdk.NewInt(10000))
	coins := sdk.NewCoins(coinEnron)
	err := suite.EnronChain.App.(*app.Enron).BankKeeper.MintCoins(suite.EnronChain.GetContext(), inflationtypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.EnronChain.App.(*app.Enron).BankKeeper.SendCoinsFromModuleToAccount(suite.EnronChain.GetContext(), inflationtypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the osmosis side which we'll use to unlock our aenron
	coinOsmo := sdk.NewCoin("uosmo", sdk.NewInt(10))
	coins = sdk.NewCoins(coinOsmo)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.MintCoins(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCOsmosisChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCOsmosisChain.GetContext(), minttypes.ModuleName, suite.IBCOsmosisChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	// Mint coins on the cosmos side which we'll use to unlock our aenron
	coinAtom := sdk.NewCoin("uatom", sdk.NewInt(10))
	coins = sdk.NewCoins(coinAtom)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.MintCoins(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, coins)
	suite.Require().NoError(err)
	err = suite.IBCCosmosChain.GetSimApp().BankKeeper.SendCoinsFromModuleToAccount(suite.IBCCosmosChain.GetContext(), minttypes.ModuleName, suite.IBCCosmosChain.SenderAccount.GetAddress(), coins)
	suite.Require().NoError(err)

	claimparams := claimtypes.DefaultParams()
	claimparams.AirdropStartTime = suite.EnronChain.GetContext().BlockTime()
	claimparams.EnableClaims = true
	suite.EnronChain.App.(*app.Enron).ClaimsKeeper.SetParams(suite.EnronChain.GetContext(), claimparams)

	params := types.DefaultParams()
	params.EnableRecovery = true
	suite.EnronChain.App.(*app.Enron).RecoveryKeeper.SetParams(suite.EnronChain.GetContext(), params)

	suite.pathOsmosisEnron = ibctesting.NewTransferPath(suite.IBCOsmosisChain, suite.EnronChain) // clientID, connectionID, channelID empty
	suite.pathCosmosEnron = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.EnronChain)
	suite.pathOsmosisCosmos = ibctesting.NewTransferPath(suite.IBCCosmosChain, suite.IBCOsmosisChain)
	suite.coordinator.Setup(suite.pathOsmosisEnron) // clientID, connectionID, channelID filled
	suite.coordinator.Setup(suite.pathCosmosEnron)
	suite.coordinator.Setup(suite.pathOsmosisCosmos)
	suite.Require().Equal("07-tendermint-0", suite.pathOsmosisEnron.EndpointA.ClientID)
	suite.Require().Equal("connection-0", suite.pathOsmosisEnron.EndpointA.ConnectionID)
	suite.Require().Equal("channel-0", suite.pathOsmosisEnron.EndpointA.ChannelID)
}

var (
	timeoutHeight = clienttypes.NewHeight(1000, 1000)

	uosmoDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uosmo",
	}

	uosmoIbcdenom = uosmoDenomtrace.IBCDenom()

	uatomDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-1",
		BaseDenom: "uatom",
	}
	uatomIbcdenom = uatomDenomtrace.IBCDenom()

	aenronDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "aenron",
	}
	aenronIbcdenom = aenronDenomtrace.IBCDenom()

	uatomOsmoDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-1",
		BaseDenom: "uatom",
	}
	uatomOsmoIbcdenom = uatomOsmoDenomtrace.IBCDenom()
)

func (suite *IBCTestingSuite) SendAndReceiveMessage(path *ibcgotesting.Path, origin *ibcgotesting.TestChain, coin string, amount int64, sender string, receiver string, seq uint64) {
	// Send coin from A to B
	transferMsg := transfertypes.NewMsgTransfer(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, sdk.NewCoin(coin, sdk.NewInt(amount)), sender, receiver, timeoutHeight, 0)
	_, err := origin.SendMsgs(transferMsg)
	suite.Require().NoError(err) // message committed
	// Recreate the packet that was sent
	transfer := transfertypes.NewFungibleTokenPacketData(coin, strconv.Itoa(int(amount)), sender, receiver)
	packet := channeltypes.NewPacket(transfer.GetBytes(), seq, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, 0)
	// Receive message on the counterparty side, and send ack
	err = path.RelayPacket(packet)
	suite.Require().NoError(err)
}

func CreatePacket(amount, denom, sender, receiver, srcPort, srcChannel, dstPort, dstChannel string, seq, timeout uint64) channeltypes.Packet {
	transfer := transfertypes.FungibleTokenPacketData{
		Amount:   amount,
		Denom:    denom,
		Receiver: sender,
		Sender:   receiver,
	}
	return channeltypes.NewPacket(
		transfer.GetBytes(),
		seq,
		srcPort,
		srcChannel,
		dstPort,
		dstChannel,
		clienttypes.ZeroHeight(), // timeout height disabled
		timeout,
	)
}
