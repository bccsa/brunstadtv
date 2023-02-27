package model

import "github.com/bcc-code/brunstadtv/backend/common"

type hasStatus interface {
	GetStatus() common.Status
}

func statusFrom(i hasStatus) Status {
	switch i.GetStatus() {
	case common.StatusUnlisted:
		return StatusUnlisted
	}
	return StatusPublished
}
