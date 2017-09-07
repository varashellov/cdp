// Code generated by cdpgen. DO NOT EDIT.

package emulation

import (
	"github.com/mafredri/cdp/rpcc"
)

// VirtualTimeBudgetExpiredClient is a client for VirtualTimeBudgetExpired events. Notification sent after the virtual time budget for the current VirtualTimePolicy has run out.
type VirtualTimeBudgetExpiredClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*VirtualTimeBudgetExpiredReply, error)
	rpcc.Stream
}

// VirtualTimeBudgetExpiredReply is the reply for VirtualTimeBudgetExpired events.
type VirtualTimeBudgetExpiredReply struct{}

// VirtualTimePausedClient is a client for VirtualTimePaused events. Notification sent after the virtual time has paused.
type VirtualTimePausedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*VirtualTimePausedReply, error)
	rpcc.Stream
}

// VirtualTimePausedReply is the reply for VirtualTimePaused events.
type VirtualTimePausedReply struct {
	VirtualTimeElapsed int `json:"virtualTimeElapsed"` // The amount of virtual time that has elapsed in milliseconds since virtual time was first enabled.
}
