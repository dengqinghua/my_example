package golang

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type base struct {
	v int
}

type Tree struct {
	// embeded struct, 我们有了 anonymous field: v
	base
	// 这里不能写成 left, right Tree
	// 但是可以为 *Tree
	left, right *Tree
}

// go test github.com/dengqinghua/golang/ -v -run TestSetImplement
// Set的简单实现 https://stackoverflow.com/a/34020023
func TestSetImplement(t *testing.T) {
	Convey("init Set", t, func() {
		names := map[string]bool{}
		abc := map[string]string{}

		treesMap := make(map[int]*Tree)

		trees := []Tree{
			Tree{base: base{v: 1}},
			Tree{base: base{v: 2}},
			Tree{base: base{v: 3}},
		}

		for i, tree := range trees {
			// treesMap[tree.v] = &tree 这种方式是不行的
			treesMap[tree.v] = &trees[i]
		}

		So(len(treesMap), ShouldEqual, 3)
		So(treesMap[2].v, ShouldEqual, 2)

		Convey("it should get", func() {
			value, ok := names["ds"]

			So(value, ShouldEqual, false)
			So(ok, ShouldBeFalse)

			v2, ok := abc["ds"]
			So(ok, ShouldBeFalse)
			So(v2, ShouldEqual, "")
		})

		Convey("it should set", func() {
			names["ds"] = true
			_, ok := names["ds"]
			value := names["ds"] // 也可以不用检查是否ok

			So(value, ShouldEqual, true)
			So(ok, ShouldBeTrue)
		})
	})
}

// go test github.com/dengqinghua/golang/ -v -run TestBasicMap
func TestBasicMap(t *testing.T) {
	Convey("TestBasicMap", t, func() {
		Convey("test nil", func() {
			Convey("when use make", func() {
				Convey("It should not equal nil", func() {
					oneMap := make(map[string]int)
					So(oneMap == nil, ShouldBeFalse)
				})
			})

			Convey("when not use make", func() {
				Convey("It should equal nil", func() {
					var oneMap map[string]int
					So(oneMap == nil, ShouldBeTrue)

					// 直接赋值是不允许的, 因为是nil
					So(func() { oneMap["ds"] = 1024 }, ShouldPanic)
				})
			})
		})

		Convey("get, put", func() {
			Convey("should get right", func() {
				oneMap := make(map[string]int)

				So(oneMap, ShouldNotBeNil)

				oneMap["dsgv"] = 587

				value, exist := oneMap["dsgv"]

				So(exist, ShouldBeTrue)
				So(value, ShouldEqual, 587)

				value, exist = oneMap["dsg"]

				So(exist, ShouldBeFalse)
				So(value, ShouldEqual, 0)
			})
		})

		Convey("init len", func() {
			Convey("should set init len", func() {
				oneMap := make(map[string]int, 10)
				oneMap["dsgv"] = 587

				So(oneMap, ShouldHaveLength, 1)
			})
		})

		Convey("set", func() {
			Convey("should init and range right", func() {
				oneMap := map[string]int{
					"dsgv": 587,
					"wjj":  523,
				}

				for name, value := range oneMap {
					if name == "dsgv" {
						So(value, ShouldEqual, 587)
					}

					if name == "wjj" {
						So(value, ShouldEqual, 523)
					}
				}
			})
		})

		Convey("init and range", func() {
			Convey("should init and range right", func() {
				oneMap := map[string]int{
					"dsgv": 587,
					"wjj":  523,
				}

				for name, value := range oneMap {
					if name == "dsgv" {
						So(value, ShouldEqual, 587)
					}

					if name == "wjj" {
						So(value, ShouldEqual, 523)
					}
				}
			})
		})
	})
}
