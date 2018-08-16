package golang

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strconv"
	"strings"
	"testing"
)

type mm struct {
	a int
	b int
}

type nn struct {
	mm
	c int
}

// go test -v -run TestNested
func TestNested(t *testing.T) {
	Convey("TestNested", t, func() {
		m := mm{1, 2}
		n := nn{m, 3}

		So(n.a, ShouldEqual, 1)
		So(n.b, ShouldEqual, 2)
		So(n.c, ShouldEqual, 3)
	})
}

// go test -v -run TestEmptyInterface
func TestEmptyInterface(t *testing.T) {
	Convey("TestEmptyInterface", t, func() {
		Convey("test equal", func() {
			a := make(map[string]interface{})

			a["errno"] = 0

			So(a["errno"] == 0, ShouldBeTrue)

			a["errno"] = "0"

			So(a["errno"] == 0, ShouldBeFalse)
		})
	})
}

// go test -v -run TestSum
func TestSum(t *testing.T) {
	Convey("Numbers", t, func() {
		// Passing Test
		Convey("Should add two numbers ", func() {
			So(Sum(1, 1), ShouldEqual, 2)
		})
	})
}

// go test -v -run TestString
func TestString(t *testing.T) {
	Convey("TestString", t, func() {
		// Passing Test
		Convey("Should like slices ", func() {
			v := "1"

			So(len(v), ShouldBeGreaterThan, 0)
			So(v[len(v)-1], ShouldEqual, 49)
		})
	})
}

// 传入的参数
// go test -v -run TestOs -args 1000 1
// Args: TestOs 1000 1
func TestOs(t *testing.T) {
	Convey("Args Test", t, func() {
		Convey("should get args", func() {
			So(len(os.Args), ShouldBeGreaterThan, 0)
			So("[]", ShouldHaveLength, 2)
		})
	})
}

const aConst = 1024

func parseOrDefault(str string, defaultVal int) int {
	b, err := strconv.Atoi(str)
	if err != nil {
		b = defaultVal
	}

	return b
}

// go test -v -run TestParseDefault
func TestParseDefault(t *testing.T) {
	Convey("ParseDefault", t, func() {
		Convey("parse with default", func() {
			var a = map[string]string{"a": "b"}
			So(parseOrDefault("1235", aConst), ShouldEqual, 1235)
			So(parseOrDefault("Ds", aConst), ShouldEqual, aConst)
			So(parseOrDefault("", aConst), ShouldEqual, aConst)

			So(parseOrDefault(a["a"], aConst), ShouldEqual, aConst)
			So(a["a"], ShouldEqual, "b")

			c, err := strconv.ParseFloat("ddd", 64)

			So(err, ShouldNotBeNil)
			So(c, ShouldEqual, 0.0)
		})
	})
}

// go test -v -run TestMultiParams
func TestMultiParams(t *testing.T) {
	Convey("TestMultiParams", t, func() {
		So(combine("1", "2", "3"), ShouldEqual, "1,2,3")
	})
}

func combine(params ...string) string {
	return strings.Join(params, ",")
}
