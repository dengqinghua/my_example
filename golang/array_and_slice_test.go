// package 的名字 和 文件夹的 名字一样
package golang

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"

	// 这里是 $GOAPTH/src/testing
	// 寻找方式为:
	// 1. ./vendor/testing
	// 2. $GOAPTH/src/testing
	// 3. $GOROOT/src/testing
	// 参考: http://lucasfcosta.com/2017/02/07/Understanding-Go-Dependency-Management.html
	"testing"
)

func extend(slice []int, element int) []int {
	n := len(slice)
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// go test github.com/dengqinghua/golang/ -v -run TestSliceExtend
func TestSliceExtend(t *testing.T) {
	Convey("TestSliceExtend", t, func() {
		var iBuffer [10]int
		panicFunc := func() {
			slice := iBuffer[0:0]
			for i := 0; i < 20; i++ {
				slice = extend(slice, i)
				fmt.Println(slice)
			}
		}

		So(panicFunc, ShouldPanic)

		// 初始化长度是3, 容量是10
		slice := make([]int, 3, 10)
		slice[0] = 0
		slice[1] = 1
		slice[2] = 2

		slice = append(slice, 3, 4, 5)
		So(slice, ShouldHaveLength, 6)
		So(cap(slice), ShouldEqual, 10)

		slice = append(slice, slice...)
		So(slice, ShouldHaveLength, 12)
		So(cap(slice), ShouldEqual, 20) // 2倍的方式进行扩张

		// ptr [0, 0, 0, 0]
		// len 4
		// cap 4
		slice = make([]int, 4)
		So(cap(slice), ShouldEqual, 4)
		So(len(slice), ShouldEqual, 4)

		// ptr [-, -, 0, 0]
		// len 2
		// cap 4
		slice = slice[:2]

		// 我理解其实这种方式的话, cap 是随着 slice 的过程动态变化的, 也是为什么叫做 slice 的原因吧
		So(len(slice), ShouldEqual, 2)
		So(cap(slice), ShouldEqual, 4)
	})
}

// One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value.
//
// go test github.com/dengqinghua/golang/ -v -run TestBasicArrayOperations
func TestBasicArrayOperations(t *testing.T) {
	Convey("init and iterate", t, func() {
		Convey("should successfully init", func() {
			var scores [10]int

			scores[0] = 19
			So(scores[0], ShouldEqual, 19)

			// 轮询操作
			for index, value := range scores {
				if index == 0 {
					So(value, ShouldEqual, 19)
				} else {
					// 默认值为0
					So(value, ShouldEqual, 0)
				}
			}
		})
	})
}

// go test github.com/dengqinghua/golang/ -v -run TestAppend
func TestAppend(t *testing.T) {
	Convey("TestAppend", t, func() {
		Convey("basic usage", func() {
			slices := make([]int, 0, 10)

			So(slices, ShouldHaveLength, 0)
			So(cap(slices), ShouldEqual, 10)

			slices = append(slices, 1, 2, 3, 4, 5, 6)

			So(len(slices), ShouldEqual, 6)
			So(cap(slices), ShouldEqual, 10)

			So(slices, ShouldResemble, []int{1, 2, 3, 4, 5, 6})

			// 奇怪的语法
			// 0:2 是指从 0 到 1, 也就是 slices[0], slices[1]
			So(slices[0:2], ShouldResemble, []int{1, 2})
			// :2 是指从 第一个元素 到 1, 也就是 slices[0], slices[1]
			So(slices[:2], ShouldResemble, []int{1, 2})
			// 3:6 是指从 3 到 5, 也就是 slices[3], slices[4], slices[5]
			So(slices[3:6], ShouldResemble, []int{4, 5, 6})
			// 3:6 是指从 3 到 最后一个元素, 也就是 slices[3], slices[4], slices[5]
			So(slices[3:], ShouldResemble, []int{4, 5, 6})

			So(slices[0], ShouldEqual, 1)
			So(slices[1], ShouldEqual, 2)

			So(slices[1], ShouldEqual, 2)

			// 现在我要删除第2个元素, slices[2] = 3
			i := 2
			// :i 代表从开始到i, i+1: 代表从i+1到最后
			slicesNew := append(slices[:i], slices[i+1:]...)

			So(slicesNew, ShouldResemble, []int{1, 2, 4, 5, 6})
			So(slices, ShouldResemble, []int{1, 2, 4, 5, 6, 6})
		})
	})
}

// 参考自 https://stackoverflow.com/a/29006008
func delFromEnd(source []int, dataToDelete int) []int {
	for i := len(source) - 1; i >= 0; i-- {
		if dataToDelete == source[i] {
			// See https://github.com/golang/go/wiki/SliceTricks#delete
			source = append(source[:i], source[i+1:]...)
		}
	}

	return source
}

func del(source []int, dataToDelete int) []int {
	var position int

	for _, data := range source {
		if data == dataToDelete {
			// See https://github.com/golang/go/wiki/SliceTricks#delete
			source = append(source[:position], source[position+1:]...)

			if position > 0 {
				position = position - 1
			}
		}

		position++
	}

	return source
}

// go test github.com/dengqinghua/golang/ -v -run TestDelete
func TestDelete(t *testing.T) {
	Convey("TestDelFromEnd", t, func() {
		Convey("basic usage", func() {
			slices := make([]int, 0, 10)

			So(slices, ShouldHaveLength, 0)
			So(cap(slices), ShouldEqual, 10)

			slices = append(slices, 1, 2, 3, 4, 5, 6)

			So(delFromEnd(slices, 4), ShouldResemble, []int{1, 2, 3, 5, 6})

			// 这个为什么 slices 为什么会变了?
			So(slices, ShouldResemble, []int{1, 2, 3, 5, 6, 6})
		})
	})

	Convey("TestDel", t, func() {
		Convey("basic usage", func() {
			slices := make([]int, 0, 10)

			So(slices, ShouldHaveLength, 0)
			So(cap(slices), ShouldEqual, 10)

			slices = append(slices, 1, 2, 3, 4, 5, 6)

			So(del(slices, 4), ShouldResemble, []int{1, 2, 3, 5, 6})
			So(slices, ShouldResemble, []int{1, 2, 3, 5, 6, 6})
		})
	})
}

// go test github.com/dengqinghua/golang/ -v -run TestNilSlice
func TestNilSlice(t *testing.T) {
	Convey("TestNilSlice", t, func() {
		Convey("use make not nil", func() {
			slices := make([]int, 0)

			So(slices == nil, ShouldBeFalse)
		})

		Convey("use int(nil) is nil", func() {
			slices := []int(nil)

			So(slices == nil, ShouldBeTrue)

			var s []int
			So(s == nil, ShouldBeTrue)

			// 设置为了nil, 注意下面的range不会报错. 不需要解决 Java 中的 NPE 问题
			s = nil

			for i := range s {
				// Never go here
				So(i, ShouldBeNil)
			}

			s = []int{}
			So(s == nil, ShouldBeFalse)

		})
	})
}

// go test github.com/dengqinghua/golang/ -v -run TestSlice
func TestSlice(t *testing.T) {
	Convey("TestSlice", t, func() {
		Convey("len and cap", func() {
			slices := make([]byte, 5, 10)
			slicesTwo := make([]byte, 8)

			Convey("should get length and capacity", func() {
				So(slices, ShouldHaveLength, 5)
				So(cap(slices), ShouldEqual, 10)

				So(slicesTwo, ShouldHaveLength, 8)
				So(cap(slicesTwo), ShouldEqual, 8)
			})
		})

		slices := make([]int, 0, 10)

		// 直接赋值是不行的, 因为没有预分配 len
		So(func() { slices[4] = 122 }, ShouldPanic)

		Convey("append", func() {
			// 可以使用append
			slices = append(slices, 10)

			Convey("should append ok", func() {
				So(slices[0], ShouldEqual, 10)
			})
		})

		Convey("re-slice", func() {
			// 可以使用append
			slicesTwo := slices[0:8]

			Convey("should append ok", func() {
				slicesTwo[7] = 1024
				So(slicesTwo[0], ShouldEqual, 0)

				// 不再 panic
				// []int{0, 0, 0, 0, 0, 0, 0, 1024}
				So(slicesTwo[7], ShouldEqual, 1024)
			})
		})
	})
}

func changeSlice(slice []int) {
	slice[0] = 100
}

// go test github.com/dengqinghua/golang/ -v -run TestPassParams
func TestPassParams(t *testing.T) {
	slice := []int{1, 2, 3}

	Convey("TestPassParams", t, func() {
		Convey("Slice", func() {
			changeSlice(slice)

			So(slice, ShouldResemble, []int{100, 2, 3})
		})

		Convey("Struct", func() {
			sStruct := s{100, 200}
			changeStruct(sStruct)

			So(sStruct.a, ShouldEqual, 100)
		})
	})
}

func changeStruct(sStruct s) {
	sStruct.a = 1
}

type s struct {
	a int
	b int
}
