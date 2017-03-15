package logx

import (
	"fmt"
	"testing"
)

func TestDebugMode(t *testing.T) {
	WithField("st", fmt.Sprintf("%d->%d", 100, 200)).Debug("Changed")
}
