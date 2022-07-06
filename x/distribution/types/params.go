package types

import (
	"fmt"

	sdk "github.com/okex/exchain/libs/cosmos-sdk/types"
	"github.com/okex/exchain/x/params"
)

const (
	// default paramspace for params keeper
	DefaultParamspace = ModuleName
)

// Parameter keys
var (
	ParamStoreKeyCommunityTax        = []byte("communitytax")
	ParamStoreKeyWithdrawAddrEnabled = []byte("withdrawaddrenabled")
	ParamStoreKeyDistributionType    = []byte("distributiontype")
)

// Params defines the set of distribution parameters.
type Params struct {
	CommunityTax        sdk.Dec `json:"community_tax" yaml:"community_tax"`
	WithdrawAddrEnabled bool    `json:"withdraw_addr_enabled" yaml:"withdraw_addr_enabled"`
	DistributionType    uint32  `json:"distribution_type" yaml:"distribution_type"`
}

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default distribution parameters
func DefaultParams() Params {
	return Params{
		CommunityTax:        sdk.NewDecWithPrec(2, 2), // 2%
		WithdrawAddrEnabled: true,
	}
}

// String returns a human readable string representation of Params
func (p Params) String() string {
	return fmt.Sprintf(`Distribution Params:
  Community Tax:          %s
  Withdraw Addr Enabled:  %t`,
		p.CommunityTax, p.WithdrawAddrEnabled)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(ParamStoreKeyCommunityTax, &p.CommunityTax, validateCommunityTax),
		params.NewParamSetPair(ParamStoreKeyWithdrawAddrEnabled, &p.WithdrawAddrEnabled, validateWithdrawAddrEnabled),
		//TODO, Check the first block of the older version is not compatible
		params.NewParamSetPair(ParamStoreKeyDistributionType, &p.DistributionType, validateDistributionType),
	}
}

// ValidateBasic performs basic validation on distribution parameters.
func (p Params) ValidateBasic() error {
	if p.CommunityTax.IsNegative() || p.CommunityTax.GT(sdk.OneDec()) {
		return fmt.Errorf(
			"community tax should non-negative and less than one: %s", p.CommunityTax,
		)
	}

	return nil
}

func validateCommunityTax(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("community tax must be not nil")
	}
	if v.IsNegative() {
		return fmt.Errorf("community tax must be positive: %s", v)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("community tax too large: %s", v)
	}

	return nil
}

func validateWithdrawAddrEnabled(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateDistributionType(i interface{}) error {
	distributionType, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if distributionType != DistributionTypeOnChain && distributionType != DistributionTypeOffChain {
		return fmt.Errorf("invalid distribution type: %d", distributionType)
	}

	return nil
}

// NewParams creates a new instance of Params
func NewParams(communityTax sdk.Dec, withdrawAddrEnabled bool, distributionType uint32) Params {
	return Params{
		CommunityTax:        communityTax,
		WithdrawAddrEnabled: withdrawAddrEnabled,
		DistributionType:    distributionType,
	}
}

// MarshalYAML implements the text format for yaml marshaling
func (p Params) MarshalYAML() (interface{}, error) {
	return p.String(), nil
}
