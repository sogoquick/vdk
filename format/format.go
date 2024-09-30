package format

import (
	"github.com/sogoquick/vdk/av/avutil"
	"github.com/sogoquick/vdk/format/aac"
	"github.com/sogoquick/vdk/format/flv"
	"github.com/sogoquick/vdk/format/mp4"
	"github.com/sogoquick/vdk/format/rtmp"
	"github.com/sogoquick/vdk/format/rtsp"
	"github.com/sogoquick/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
