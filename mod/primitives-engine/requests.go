// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package engineprimitives

import (
	"unsafe"

	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	gengine "github.com/ethereum/go-ethereum/beacon/engine"
	coretypes "github.com/ethereum/go-ethereum/core/types"
)

// NewPayloadRequest as per the Ethereum 2.0 specification:
// https://github.com/ethereum/consensus-specs/blob/dev/specs/deneb/beacon-chain.md#modified-newpayloadrequest
//
//nolint:lll
type NewPayloadRequest struct {
	// ExecutionPayload is the payload to the execution client.
	ExecutionPayload ExecutionPayload
	// VersionedHashes is the versioned hashes of the execution payload.
	VersionedHashes []common.ExecutionHash
	// ParentBeaconBlockRoot is the root of the parent beacon block.
	ParentBeaconBlockRoot *primitives.Root
	// SkipIfExists is a flag that indicates if the payload should be skipped
	// if it already exists in the database of the execution client.
	SkipIfExists bool
	// Optimistic is a flag that indicates if the payload should be
	// optimistically deemed valid. This is useful during syncing.
	Optimistic bool
}

// BuildNewPayloadRequest builds a new payload request.
func BuildNewPayloadRequest(
	executionPayload ExecutionPayload,
	versionedHashes []common.ExecutionHash,
	parentBeaconBlockRoot *primitives.Root,
	skipIfExists bool,
	optimistic bool,
) *NewPayloadRequest {
	return &NewPayloadRequest{
		ExecutionPayload:      executionPayload,
		VersionedHashes:       versionedHashes,
		ParentBeaconBlockRoot: parentBeaconBlockRoot,
		SkipIfExists:          skipIfExists,
		Optimistic:            optimistic,
	}
}

// HasHValidVersionAndBlockHashes checks if the version and block hashes are
// valid.
// As per the Ethereum 2.0 specification:
// https://github.com/ethereum/consensus-specs/blob/v1.4.0-beta.2/specs/deneb/beacon-chain.md#is_valid_block_hash
// https://github.com/ethereum/consensus-specs/blob/v1.4.0-beta.2/specs/deneb/beacon-chain.md#is_valid_versioned_hashes
//
//nolint:lll
func (n *NewPayloadRequest) HasValidVersionedAndBlockHashes() error {
	payload := n.ExecutionPayload
	withdrawals := payload.GetWithdrawals()
	data := gengine.ExecutableData{
		ParentHash:    payload.GetParentHash(),
		FeeRecipient:  payload.GetFeeRecipient(),
		StateRoot:     common.ExecutionHash(payload.GetStateRoot()),
		ReceiptsRoot:  common.ExecutionHash(payload.GetReceiptsRoot()),
		LogsBloom:     payload.GetLogsBloom(),
		Random:        common.ExecutionHash(payload.GetPrevRandao()),
		Number:        payload.GetNumber().Unwrap(),
		GasLimit:      payload.GetGasLimit().Unwrap(),
		GasUsed:       payload.GetGasUsed().Unwrap(),
		Timestamp:     payload.GetTimestamp().Unwrap(),
		ExtraData:     payload.GetExtraData(),
		BaseFeePerGas: payload.GetBaseFeePerGas().UnwrapBig(),
		BlockHash:     payload.GetBlockHash(),
		Transactions:  payload.GetTransactions(),
		//#nosec:G103 // henlo I am the captain now.
		Withdrawals:   *(*[]*coretypes.Withdrawal)(unsafe.Pointer(&withdrawals)),
		BlobGasUsed:   payload.GetBlobGasUsed().UnwrapPtr(),
		ExcessBlobGas: payload.GetExcessBlobGas().UnwrapPtr(),
	}
	_, err := gengine.ExecutableDataToBlock(
		data,
		n.VersionedHashes,
		(*common.ExecutionHash)(n.ParentBeaconBlockRoot),
	)
	return err
}

// ForkchoiceUpdateRequest.
type ForkchoiceUpdateRequest struct {
	// State is the forkchoice state.
	State *ForkchoiceState
	// PayloadAttributes is the payload attributer.
	PayloadAttributes PayloadAttributer
	// ForkVersion is the fork version that we
	// are going to be submitting for.
	ForkVersion uint32
}

// BuildForkchoiceUpdateRequest builds a forkchoice update request.
func BuildForkchoiceUpdateRequest(
	state *ForkchoiceState,
	payloadAttributes PayloadAttributer,
	forkVersion uint32,
) *ForkchoiceUpdateRequest {
	return &ForkchoiceUpdateRequest{
		State:             state,
		PayloadAttributes: payloadAttributes,
		ForkVersion:       forkVersion,
	}
}

// GetPayloadRequest represents a request to get a payload.
type GetPayloadRequest struct {
	// PayloadID is the payload ID.
	PayloadID PayloadID
	// ForkVersion is the fork version that we are
	// currently on.
	ForkVersion uint32
}

// BuildGetPayloadRequest builds a get payload request.
func BuildGetPayloadRequest(
	payloadID PayloadID,
	forkVersion uint32,
) *GetPayloadRequest {
	return &GetPayloadRequest{
		PayloadID:   payloadID,
		ForkVersion: forkVersion,
	}
}