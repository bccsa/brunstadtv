package notifications

import (
	"context"

	"github.com/bcc-code/bcc-media-platform/backend/common"
	"github.com/bcc-code/bcc-media-platform/backend/targets"
	"github.com/bcc-code/mediabank-bridge/log"
	"github.com/google/uuid"
)

// ResolveTargets resolves targetIDs to device tokens
func (u *Utils) ResolveTargets(ctx context.Context, targetIDs []uuid.UUID) ([]common.Device, error) {
	log.L.Debug().Int("targetCount", len(targetIDs)).Msg("Resolving targets")
	targetRows, err := u.queries.GetTargets(ctx, targetIDs)
	if err != nil {
		return nil, err
	}

	var devices []common.Device
	for _, t := range targetRows {
		target := common.Target(t)
		ds, err := targets.ResolveDevices(ctx, u.queries, target)
		if err != nil {
			return nil, err
		}
		log.L.Debug().Int("deviceCount", len(ds)).Msg("Resolved target, retrieved devices")
		devices = append(devices, ds...)
	}
	return devices, nil
}
