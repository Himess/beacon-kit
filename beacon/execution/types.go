// SPDX-License-Identifier: MIT
//
// Copyright (c) 2023 Berachain Foundation
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

package execution

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/itsdevbear/bolaris/async/dispatch"
	"github.com/itsdevbear/bolaris/types/state"
	"github.com/prysmaticlabs/prysm/v4/consensus-types/primitives"
)

// BeaconStateProvider is an interface that wraps the basic BeaconState method.
type BeaconStateProvider interface {
	// BeaconState provides access to the underlying beacon state.
	BeaconState(ctx context.Context) state.BeaconState
}

// GrandCentralDispatch is an interface that wraps the basic GetQueue method.
// It is used to retrieve a dispatch queue by its ID.
type GrandCentralDispatch interface {
	// GetQueue returns a queue with the provided ID.
	GetQueue(id string) dispatch.Queue
}

// NotifyForkchoiceUpdateArg is the argument for the forkchoice
// update notification `notifyForkchoiceUpdate`.
type NotifyForkchoiceUpdateArg struct {
	// headHash is the hash of the head block we are building ontop of.=
	headHash common.Hash
	// safeHash is the hash of the last safe block.
	safeHash common.Hash
	// finalHash is the hash of the last finalized block.
	finalHash common.Hash
}

type FCUConfig struct {
	HeadEth1Hash  common.Hash
	ProposingSlot primitives.Slot
	// attributes    payloadattribute.Attributer
}

// NewNotifyForkchoiceUpdateArg creates a new NotifyForkchoiceUpdateArg.
func NewNotifyForkchoiceUpdateArg(
	headHash common.Hash,
	safeHash common.Hash,
	finalHash common.Hash,
) *NotifyForkchoiceUpdateArg {
	return &NotifyForkchoiceUpdateArg{
		headHash:  headHash,
		safeHash:  safeHash,
		finalHash: finalHash,
	}
}
