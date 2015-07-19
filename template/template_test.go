package template

import (
	//"bytes"
	"testing"

	//"github.com/jmervine/getdown/config"

	. "gopkg.in/jmervine/GoT.v1"
)

const TEMPLATE = "../template.html"

func TestNewTemplate(T *testing.T) {
	//var buf bytes.Buffer
	t, e := NewTemplate(TEMPLATE)
	Go(T).AssertNil(e)
	Go(T).RefuteNil(t)
}
