package codenames

import (
	"io"
)

type discardStore struct{}

func (ds discardStore) Save(*Game) error           { return nil }
func (ds discardStore) Delete(*Game) error         { return nil }
func (ds discardStore) Checkpoint(io.Writer) error { return nil }
