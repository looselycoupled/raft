package raft

import "errors"

// Standard errors for primary operations.
var (
	ErrCommitIndex      = errors.New("commit index does not refer to an entry in the log")
	ErrAlreadyCommitted = errors.New("commit index precedes current commit index")
	ErrMissingCommit    = errors.New("cannot commit entry higher than found in log")
	ErrNotImplemented   = errors.New("functionality not implemented yet")
	ErrEventTypeError   = errors.New("captured event with wrong value type")
	ErrEventSourceError = errors.New("captured event with wrong source type")
	ErrUnknownState     = errors.New("raft in an unknown state")
	ErrNotListening     = errors.New("replica is not listening for events")
)
