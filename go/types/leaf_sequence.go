// Copyright 2016 Attic Labs, Inc. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

type leafSequence struct {
	vr     ValueReader
	length int
	kind   NomsKind
}

func (seq leafSequence) seqLen() int {
	return seq.length
}

func (seq leafSequence) numLeaves() uint64 {
	return uint64(seq.length)
}

func (seq leafSequence) valueReader() ValueReader {
	return seq.vr
}

func (seq leafSequence) getChildSequence(idx int) sequence {
	return nil
}

func (seq leafSequence) Kind() NomsKind {
	return seq.kind
}
