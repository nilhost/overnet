package json

//go:generate go run github.com/nilhost/overnet/common/errors/errorgen

import (
	"io"

	"github.com/nilhost/overnet"
	"github.com/nilhost/overnet/common"
	"github.com/nilhost/overnet/common/cmdarg"
	"github.com/nilhost/overnet/infra/conf/serial"
	"github.com/nilhost/overnet/main/confloader"
)

func init() {
	common.Must(core.RegisterConfigLoader(&core.ConfigFormat{
		Name:      "JSON",
		Extension: []string{"json"},
		Loader: func(input interface{}) (*core.Config, error) {
			switch v := input.(type) {
			case cmdarg.Arg:
				r, err := confloader.LoadExtConfig(v)
				if err != nil {
					return nil, newError("failed to execute v2ctl to convert config file.").Base(err).AtWarning()
				}
				return core.LoadConfig("protobuf", "", r)
			case io.Reader:
				return serial.LoadJSONConfig(v)
			default:
				return nil, newError("unknow type")
			}
		},
	}))
}
