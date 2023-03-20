// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package t8ntool

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/aroeum-network/go-aroeum/common"
	"github.com/aroeum-network/go-aroeum/common/math"
	"github.com/aroeum-network/go-aroeum/core/types"
)

var _ = (*stEnvMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (s stEnv) MarshalJSON() ([]byte, error) {
	type stEnv struct {
		Coinbase         common.UnprefixedAddress            `json:"currentCoinbase"   gencodec:"required"`
		Difficulty       *math.HexOrDecimal256               `json:"currentDifficulty"`
		Random           *math.HexOrDecimal256               `json:"currentRandom"`
		ParentDifficulty *math.HexOrDecimal256               `json:"parentDifficulty"`
		ParentBaseFee    *math.HexOrDecimal256               `json:"parentBaseFee,omitempty"`
		ParentGasUsed    math.HexOrDecimal64                 `json:"parentGasUsed,omitempty"`
		ParentGasLimit   math.HexOrDecimal64                 `json:"parentGasLimit,omitempty"`
		GasLimit         math.HexOrDecimal64                 `json:"currentGasLimit"   gencodec:"required"`
		Number           math.HexOrDecimal64                 `json:"currentNumber"     gencodec:"required"`
		Timestamp        math.HexOrDecimal64                 `json:"currentTimestamp"  gencodec:"required"`
		ParentTimestamp  math.HexOrDecimal64                 `json:"parentTimestamp,omitempty"`
		BlockHashes      map[math.HexOrDecimal64]common.Hash `json:"blockHashes,omitempty"`
		Ommers           []ommer                             `json:"ommers,omitempty"`
		Withdrawals      []*types.Withdrawal                 `json:"withdrawals,omitempty"`
		BaseFee          *math.HexOrDecimal256               `json:"currentBaseFee,omitempty"`
		ParentUncleHash  common.Hash                         `json:"parentUncleHash"`
	}
	var enc stEnv
	enc.Coinbase = common.UnprefixedAddress(s.Coinbase)
	enc.Difficulty = (*math.HexOrDecimal256)(s.Difficulty)
	enc.Random = (*math.HexOrDecimal256)(s.Random)
	enc.ParentDifficulty = (*math.HexOrDecimal256)(s.ParentDifficulty)
	enc.ParentBaseFee = (*math.HexOrDecimal256)(s.ParentBaseFee)
	enc.ParentGasUsed = math.HexOrDecimal64(s.ParentGasUsed)
	enc.ParentGasLimit = math.HexOrDecimal64(s.ParentGasLimit)
	enc.GasLimit = math.HexOrDecimal64(s.GasLimit)
	enc.Number = math.HexOrDecimal64(s.Number)
	enc.Timestamp = math.HexOrDecimal64(s.Timestamp)
	enc.ParentTimestamp = math.HexOrDecimal64(s.ParentTimestamp)
	enc.BlockHashes = s.BlockHashes
	enc.Ommers = s.Ommers
	enc.Withdrawals = s.Withdrawals
	enc.BaseFee = (*math.HexOrDecimal256)(s.BaseFee)
	enc.ParentUncleHash = s.ParentUncleHash
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (s *stEnv) UnmarshalJSON(input []byte) error {
	type stEnv struct {
		Coinbase         *common.UnprefixedAddress           `json:"currentCoinbase"   gencodec:"required"`
		Difficulty       *math.HexOrDecimal256               `json:"currentDifficulty"`
		Random           *math.HexOrDecimal256               `json:"currentRandom"`
		ParentDifficulty *math.HexOrDecimal256               `json:"parentDifficulty"`
		ParentBaseFee    *math.HexOrDecimal256               `json:"parentBaseFee,omitempty"`
		ParentGasUsed    *math.HexOrDecimal64                `json:"parentGasUsed,omitempty"`
		ParentGasLimit   *math.HexOrDecimal64                `json:"parentGasLimit,omitempty"`
		GasLimit         *math.HexOrDecimal64                `json:"currentGasLimit"   gencodec:"required"`
		Number           *math.HexOrDecimal64                `json:"currentNumber"     gencodec:"required"`
		Timestamp        *math.HexOrDecimal64                `json:"currentTimestamp"  gencodec:"required"`
		ParentTimestamp  *math.HexOrDecimal64                `json:"parentTimestamp,omitempty"`
		BlockHashes      map[math.HexOrDecimal64]common.Hash `json:"blockHashes,omitempty"`
		Ommers           []ommer                             `json:"ommers,omitempty"`
		Withdrawals      []*types.Withdrawal                 `json:"withdrawals,omitempty"`
		BaseFee          *math.HexOrDecimal256               `json:"currentBaseFee,omitempty"`
		ParentUncleHash  *common.Hash                        `json:"parentUncleHash"`
	}
	var dec stEnv
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Coinbase == nil {
		return errors.New("missing required field 'currentCoinbase' for stEnv")
	}
	s.Coinbase = common.Address(*dec.Coinbase)
	if dec.Difficulty != nil {
		s.Difficulty = (*big.Int)(dec.Difficulty)
	}
	if dec.Random != nil {
		s.Random = (*big.Int)(dec.Random)
	}
	if dec.ParentDifficulty != nil {
		s.ParentDifficulty = (*big.Int)(dec.ParentDifficulty)
	}
	if dec.ParentBaseFee != nil {
		s.ParentBaseFee = (*big.Int)(dec.ParentBaseFee)
	}
	if dec.ParentGasUsed != nil {
		s.ParentGasUsed = uint64(*dec.ParentGasUsed)
	}
	if dec.ParentGasLimit != nil {
		s.ParentGasLimit = uint64(*dec.ParentGasLimit)
	}
	if dec.GasLimit == nil {
		return errors.New("missing required field 'currentGasLimit' for stEnv")
	}
	s.GasLimit = uint64(*dec.GasLimit)
	if dec.Number == nil {
		return errors.New("missing required field 'currentNumber' for stEnv")
	}
	s.Number = uint64(*dec.Number)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'currentTimestamp' for stEnv")
	}
	s.Timestamp = uint64(*dec.Timestamp)
	if dec.ParentTimestamp != nil {
		s.ParentTimestamp = uint64(*dec.ParentTimestamp)
	}
	if dec.BlockHashes != nil {
		s.BlockHashes = dec.BlockHashes
	}
	if dec.Ommers != nil {
		s.Ommers = dec.Ommers
	}
	if dec.Withdrawals != nil {
		s.Withdrawals = dec.Withdrawals
	}
	if dec.BaseFee != nil {
		s.BaseFee = (*big.Int)(dec.BaseFee)
	}
	if dec.ParentUncleHash != nil {
		s.ParentUncleHash = *dec.ParentUncleHash
	}
	return nil
}
