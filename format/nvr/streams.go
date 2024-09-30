package nvr

import "github.com/sogoquick/vdk/av"

type Stream struct {
	codec av.CodecData
	idx   int
}
