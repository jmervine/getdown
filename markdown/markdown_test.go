package markdown

import (
	. "github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
	"testing"

	"github.com/jmervine/getdown/config"
)

const FILE = "../README.md"

func TestNew(T *testing.T) {
	md, err := New(FILE, &config.Default)

	Go(T).AssertNil(err)
	Go(T).RefuteNil(md)
}
