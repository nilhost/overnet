package ctlcmd

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/nilhost/overnet/common/buf"
	"github.com/nilhost/overnet/common/platform"
)

//go:generate go run github.com/nilhost/overnet/common/errors/errorgen

func Run(args []string, input io.Reader) (buf.MultiBuffer, error) {
	overctl := platform.GetToolLocation("overctl")
	if _, err := os.Stat(overctl); err != nil {
		return nil, newError("v2ctl doesn't exist").Base(err)
	}

	var errBuffer buf.MultiBufferContainer
	var outBuffer buf.MultiBufferContainer

	cmd := exec.Command(overctl, args...)
	cmd.Stderr = &errBuffer
	cmd.Stdout = &outBuffer
	cmd.SysProcAttr = getSysProcAttr()
	if input != nil {
		cmd.Stdin = input
	}

	if err := cmd.Start(); err != nil {
		return nil, newError("failed to start overctl").Base(err)
	}

	if err := cmd.Wait(); err != nil {
		msg := "failed to execute overctl"
		if errBuffer.Len() > 0 {
			msg += ": \n" + strings.TrimSpace(errBuffer.MultiBuffer.String())
		}
		return nil, newError(msg).Base(err)
	}

	// log stderr, info message
	if !errBuffer.IsEmpty() {
		newError("<overctl message> \n", strings.TrimSpace(errBuffer.MultiBuffer.String())).AtInfo().WriteToLog()
	}

	return outBuffer.MultiBuffer, nil
}
