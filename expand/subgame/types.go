package subgame

import (
	"github.com/JFJun/go-substrate-rpc-client/v3/types"
	"github.com/SubGame-Network/bifrost-go/expand/base"
)

type SubGameEventRecords struct {
	base.BaseEventRecords
	Claims_Claimed                    []EventClaimsClaimed
	ElectionsPhragmen_VoterReported   []EventElectionsPhragmenVoterReported
	ElectionsPhragmen_MemberRenounced []EventElectionsPhragmenMemberRenounced
	ElectionsPhragmen_MemberKicked    []EventElectionsPhragmenMemberKicked
	ElectionsPhragmen_ElectionError   []EventElectionsPhragmenElectionError
	ElectionsPhragmen_EmptyTerm       []EventElectionsPhragmenEmptyTerm
	//ElectionsPhragmen_NewTerm		[]EventElectionsPhragmenNewTerm		暂不支持解析
	Democracy_Blacklisted []EventDemocracyBlacklisted
	//2021-06-09
	Staking_Kicked                                  []EventStakingKicked
	ElectionsPhragmen_NewTerm                       []EventElectionsPhragmenNewTerm
	ElectionsPhragmen_Renounced                     []EventElectionsPhragmenRenounced
	ElectionsPhragmen_CandidateSlashed              []EventElectionsPhragmenCandidateSlashed
	ElectionsPhragmen_SeatHolderSlashed             []EventElectionsPhragmenSeatHolderSlashed
	Bounties_BountyProposed                         []EventBountiesBountyProposed
	Bounties_BountyRejected                         []EventBountiesBountyRejected
	Bounties_BountyBecameActive                     []EventBountiesBountyBecameActive
	Bounties_BountyAwarded                          []EventBountiesBountyAwarded
	Bounties_BountyClaimed                          []EventBountiesBountyClaimed
	Bounties_BountyCanceled                         []EventBountiesBountyCanceled
	Bounties_BountyExtended                         []EventBountiesBountyExtended
	Tips_NewTip                                     []EventTipsNewTip
	Tips_TipClosing                                 []EventTipsTipClosing
	Tips_TipClosed                                  []EventTipsTipClosed
	Tips_TipRetracted                               []EventTipsTipRetracted
	Tips_TipSlashed                                 []EventTipsTipSlashed
	ElectionProviderMultiPhase_SolutionStored       []EventElectionProviderMultiPhaseSolutionStored
	ElectionProviderMultiPhase_ElectionFinalized    []EventElectionProviderMultiPhaseElectionFinalized
	ElectionProviderMultiPhase_Rewarded             []EventElectionProviderMultiPhaseRewarded
	ElectionProviderMultiPhase_Slashed              []EventElectionProviderMultiPhaseSlashed
	ElectionProviderMultiPhase_SignedPhaseStarted   []EventElectionProviderMultiPhaseSignedPhaseStarted
	ElectionProviderMultiPhase_UnsignedPhaseStarted []EventElectionProviderMultiPhaseUnsignedPhaseStarted

	Chips_BuyChips           []Chips_BuyChips
	Chips_Redemption         []Chips_Redemption
	Chips_Reserve            []Chips_Reserve
	Chips_Unreserve          []Chips_Unreserve
	Chips_RepatriateReserved []Chips_RepatriateReserved

	Bridge_ReceiveBridge []Bridge_ReceiveBridge
	Bridge_Send          []Bridge_Send

	Stake_SignUp   []Stake_SignUp
	Stake_Stake    []Stake_Stake
	Stake_Unlock   []Stake_Unlock
	Stake_Withdraw []Stake_Withdraw

	SubgameAssets_Created           []SubgameAssets_Created
	SubgameAssets_Issued            []SubgameAssets_Issued
	SubgameAssets_Transferred       []SubgameAssets_Transferred
	SubgameAssets_Burned            []SubgameAssets_Burned
	SubgameAssets_TeamChanged       []SubgameAssets_TeamChanged
	SubgameAssets_OwnerChanged      []SubgameAssets_OwnerChanged
	SubgameAssets_ForceTransferred  []SubgameAssets_ForceTransferred
	SubgameAssets_Frozen            []SubgameAssets_Frozen
	SubgameAssets_Thawed            []SubgameAssets_Thawed
	SubgameAssets_AssetFrozen       []SubgameAssets_AssetFrozen
	SubgameAssets_AssetThawed       []SubgameAssets_AssetThawed
	SubgameAssets_Destroyed         []SubgameAssets_Destroyed
	SubgameAssets_ForceCreated      []SubgameAssets_ForceCreated
	SubgameAssets_MaxZombiesChanged []SubgameAssets_MaxZombiesChanged
	SubgameAssets_MetadataSet       []SubgameAssets_MetadataSet
}

type EventElectionProviderMultiPhaseUnsignedPhaseStarted struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}
type EventElectionProviderMultiPhaseSignedPhaseStarted struct {
	Phase  types.Phase
	U32    types.U32
	Topics []types.Hash
}
type EventElectionProviderMultiPhaseSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionProviderMultiPhaseRewarded struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionProviderMultiPhaseElectionFinalized struct {
	Phase                  types.Phase
	Option_ElectionCompute base.OptionElectionCompute
	Topics                 []types.Hash
}
type EventElectionProviderMultiPhaseSolutionStored struct {
	Phase           types.Phase
	ElectionCompute types.ElectionCompute
	Topics          []types.Hash
}
type EventTipsTipSlashed struct {
	Phase     types.Phase
	Hash      types.Hash
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventTipsTipRetracted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventTipsTipClosed struct {
	Phase     types.Phase
	Hash      types.Hash
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventTipsTipClosing struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventTipsNewTip struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}
type EventBountiesBountyExtended struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventBountiesBountyCanceled struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventBountiesBountyClaimed struct {
	Phase       types.Phase
	BountyIndex types.U32
	Balance     types.U128
	AccountId   types.AccountID
	Topics      []types.Hash
}
type EventBountiesBountyAwarded struct {
	Phase       types.Phase
	BountyIndex types.U32
	AccountId   types.AccountID
	Topics      []types.Hash
}
type EventBountiesBountyBecameActive struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventBountiesBountyRejected struct {
	Phase       types.Phase
	BountyIndex types.U32
	Balance     types.U128
	Topics      []types.Hash
}
type EventBountiesBountyProposed struct {
	Phase       types.Phase
	BountyIndex types.U32
	Topics      []types.Hash
}
type EventElectionsPhragmenSeatHolderSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventElectionsPhragmenCandidateSlashed struct {
	Phase     types.Phase
	AccountId types.AccountID
	Balance   types.U128
	Topics    []types.Hash
}
type EventElectionsPhragmenRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenNewTerm struct {
	Phase  types.Phase
	Abs    []base.AccountBalance
	Topics []types.Hash
}
type EventStakingKicked struct {
	Phase      types.Phase
	AccountId1 types.AccountID
	AccountId2 types.AccountID
	Topics     []types.Hash
}

func (p SubGameEventRecords) GetBalancesTransfer() []types.EventBalancesTransfer {
	return p.Balances_Transfer
}

func (p SubGameEventRecords) GetSystemExtrinsicSuccess() []types.EventSystemExtrinsicSuccess {
	return p.System_ExtrinsicSuccess
}

func (p SubGameEventRecords) GetSystemExtrinsicFailed() []types.EventSystemExtrinsicFailed {
	return p.System_ExtrinsicFailed
}

type EventDemocracyBlacklisted struct {
	Phase  types.Phase
	Hash   types.Hash
	Topics []types.Hash
}

//type EventElectionsPhragmenNewTerm struct {
//	Phase    types.Phase
//	Vec
//	Topics []types.Hash
//}
type EventElectionsPhragmenEmptyTerm struct {
	Phase types.Phase

	Topics []types.Hash
}
type EventElectionsPhragmenElectionError struct {
	Phase  types.Phase
	Topics []types.Hash
}
type EventElectionsPhragmenMemberKicked struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenMemberRenounced struct {
	Phase     types.Phase
	AccountId types.AccountID
	Topics    []types.Hash
}
type EventElectionsPhragmenVoterReported struct {
	Phase  types.Phase
	Who1   types.AccountID
	Who2   types.AccountID
	Bool   types.Bool
	Topics []types.Hash
}
type EventClaimsClaimed struct {
	Phase           types.Phase
	AccountId       types.AccountID
	EthereumAddress base.VecU8L20
	Balance         types.U128
	Topics          []types.Hash
}

type Chips_BuyChips struct {
	Phase  types.Phase
	From   types.AccountID
	Value  types.U128
	Topics []types.Hash
}
type Chips_Redemption struct {
	Phase  types.Phase
	From   types.AccountID
	Value  types.U128
	Topics []types.Hash
}
type Chips_Reserve struct {
	Phase  types.Phase
	From   types.AccountID
	Value  types.U128
	Topics []types.Hash
}
type Chips_Unreserve struct {
	Phase  types.Phase
	From   types.AccountID
	Value  types.U128
	Topics []types.Hash
}
type Chips_RepatriateReserved struct {
	Phase  types.Phase
	From   types.AccountID
	To     types.AccountID
	Value  types.U128
	Topics []types.Hash
}

type Bridge_ReceiveBridge struct {
	Phase       types.Phase
	From        types.AccountID
	ToAddress   types.Text
	ToChainType types.U8
	ToCoinType  types.U8
	Amount      types.U128
	Topics      []types.Hash
}
type Bridge_Send struct {
	Phase         types.Phase
	To            types.AccountID
	Amount        types.U128
	ReceiveTxHash types.Text
	Topics        []types.Hash
}

type Stake_SignUp struct {
	Phase           types.Phase
	From            types.AccountID
	Account         types.Text
	ReferrerAccount types.Text
	Topics          []types.Hash
}
type Stake_Stake struct {
	Phase  types.Phase
	From   types.AccountID
	Amount types.U128
	Topics []types.Hash
}
type Stake_Unlock struct {
	Phase  types.Phase
	From   types.AccountID
	Amount types.U128
	Topics []types.Hash
}
type Stake_Withdraw struct {
	Phase  types.Phase
	From   types.AccountID
	Amount types.U128
	Topics []types.Hash
}
type SubgameAssets_Created struct {
	Phase   types.Phase
	AssetId types.U32
	Owner   types.AccountID
	Admin   types.AccountID
	Topics  []types.Hash
}
type SubgameAssets_Issued struct {
	Phase       types.Phase
	AssetId     types.U32
	Beneficiary types.AccountID
	Amount      types.U64
	Topics      []types.Hash
}
type SubgameAssets_Transferred struct {
	Phase   types.Phase
	AssetId types.U32
	Origin  types.AccountID
	Dest    types.AccountID
	Amount  types.U64
	Topics  []types.Hash
}
type SubgameAssets_Burned struct {
	Phase   types.Phase
	AssetId types.U32
	Who     types.AccountID
	Burned  types.U64
	Topics  []types.Hash
}
type SubgameAssets_TeamChanged struct {
	Phase   types.Phase
	AssetId types.U32
	Issuer  types.AccountID
	Admin   types.AccountID
	Freezer types.AccountID
	Topics  []types.Hash
}
type SubgameAssets_OwnerChanged struct {
	Phase   types.Phase
	AssetId types.U32
	Owner   types.AccountID
	Topics  []types.Hash
}
type SubgameAssets_ForceTransferred struct {
	Phase   types.Phase
	AssetId types.U32
	source  types.AccountID
	dest    types.AccountID
	amount  types.U64
	Topics  []types.Hash
}
type SubgameAssets_Frozen struct {
	Phase   types.Phase
	AssetId types.U32
	Who     types.AccountID
	Topics  []types.Hash
}
type SubgameAssets_Thawed struct {
	Phase   types.Phase
	AssetId types.U32
	Who     types.AccountID
	Topics  []types.Hash
}
type SubgameAssets_AssetFrozen struct {
	Phase   types.Phase
	AssetId types.U32
	Topics  []types.Hash
}
type SubgameAssets_AssetThawed struct {
	Phase   types.Phase
	AssetId types.U32
	Topics  []types.Hash
}
type SubgameAssets_Destroyed struct {
	Phase   types.Phase
	AssetId types.U32
	Topics  []types.Hash
}
type SubgameAssets_ForceCreated struct {
	Phase   types.Phase
	AssetId types.U32
	Owner   types.AccountID
	Topics  []types.Hash
}
type SubgameAssets_MaxZombiesChanged struct {
	Phase      types.Phase
	AssetId    types.U32
	MaxZombies types.U32
	Topics     []types.Hash
}
type SubgameAssets_MetadataSet struct {
	Phase    types.Phase
	AssetId  types.U32
	Name     types.Text
	Symbol   types.Text
	Decimals types.U8
	Topics   []types.Hash
}
