// +build debug

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	"github.com/filecoin-project/specs-actors/actors/builtin/miner"
	miner2 "github.com/filecoin-project/specs-actors/v2/actors/builtin/miner"
	"os"
)

const BootstrappersFile = ""
const GenesisFile = ""

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1
const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3
const UpgradeTapeHeight = -4

const UpgradeActorsV2Height = 10
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 15
const UpgradeCalicoHeight = 20
const UpgradePersianHeight = 25
const UpgradeOrangeHeight = 27
const UpgradeClausHeight = 30

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

func init() {
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg2KiBV1,
		abi.RegisteredSealProof_StackedDrg8MiBV1,
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))

	miner.PreCommitChallengeDelay = abi.ChainEpoch(10)
	miner2.PreCommitChallengeDelay = abi.ChainEpoch(10)

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	Devnet = false
}

const BlockDelaySecs = uint64(20)

const PropagationDelaySecs = uint64(5)

// SlashablePowerDelay is the number of epochs after ElectionPeriodStart, after
// which the miner is slashed
//
// Epochs
const SlashablePowerDelay = 20

// Epochs
const InteractivePoRepConfidence = 6

const BootstrapPeerThreshold = 1
