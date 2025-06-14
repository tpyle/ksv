package clipboard

import (
	"github.com/tpyle/ksv/lib/ksverrors"
	"golang.design/x/clipboard"
)

var (
	canUseClipboard = true
)

func init() {
	err := clipboard.Init()
	if err != nil {
		canUseClipboard = false
	}
}

func Copy(text string) error {
	if !canUseClipboard {
		return ksverrors.ErrClipboardUnavailable
	}
	_ = clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}
