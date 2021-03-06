package golo

import (
	"github.com/gofrs/uuid"
)

const (
	// DefaultSequenceID is a uuid which will be returned should uuid.NewV4
	// fail. It can be safely compared with whatever is returned from
	// loadtest.SequenceID()- this uuid is a v5 uuid in the DNS namespace
	// whereas SequenceID() returns a v4 uuid.
	// see script/make_uuid.go in source repo for more information.
	DefaultSequenceID = "c276c8c7-6fec-5aa9-b6bd-4de12a49a9bb"
)

// NewSequenceID will return a fresh v4 uuid for sequences
// of requests to use, to allow for ease of grouping journeys
// together. This function swallows errors; should an error occur
// then this will, instead, return loadtest.DefaultSequenceID.
// Thus: a usable ID can always be guaranteed from this function
func NewSequenceID() string {
	s, err := uuid.NewV4()
	if err != nil {
		return DefaultSequenceID
	}

	return s.String()
}
