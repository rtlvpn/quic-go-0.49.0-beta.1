package wire

import "github.com/sagernet/quic-go/internal/protocol"

// A PaddingFrame is a padding frame
type PaddingFrame struct {
	PaddingLength protocol.ByteCount
}

func (f *PaddingFrame) Append(b []byte, _ protocol.Version) ([]byte, error) {
	return append(b, make([]byte, f.PaddingLength)...), nil
}

// Length returns the length of a written frame
func (f *PaddingFrame) Length(_ protocol.Version) protocol.ByteCount {
	return f.PaddingLength
}
