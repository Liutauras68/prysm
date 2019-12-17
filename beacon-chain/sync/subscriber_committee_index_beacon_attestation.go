package sync

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
)

func (r *Service) committeeIndexBeaconAttestationSubscriber(ctx context.Context, msg proto.Message) error {

	return nil
}

func (r *Service) currentCommitteeIndex() int {
	state, err := r.chain.HeadState(context.Background())
	if err != nil {
		panic(err)
	}
	if state == nil {
		return 0
	}
	count, err := helpers.CommitteeCountAtSlot(state, helpers.StartSlot(helpers.CurrentEpoch(state)))
	if err != nil {
		panic(err)
	}
	return int(count)
}