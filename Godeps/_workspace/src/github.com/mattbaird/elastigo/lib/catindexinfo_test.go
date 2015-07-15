package elastigo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCatIndexInfo(t *testing.T) {
	Convey("Create index line from a broken index listing", t, func() {
		_, err := NewCatIndexInfo("red ")
		So(err, ShouldNotBeNil)
	})
	Convey("Create index line from a bad shards index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar a 1 1234 3 11000 13000")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 0)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 1234)
		So(i.Docs.Deleted, ShouldEqual, 3)
		So(i.Store.Size, ShouldEqual, 11000)
		So(i.Store.PriSize, ShouldEqual, 13000)
	})
	Convey("Create index line from a bad replicas index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 a 1234 3 11000 13000")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 0)
		So(i.Docs.Count, ShouldEqual, 1234)
		So(i.Docs.Deleted, ShouldEqual, 3)
		So(i.Store.Size, ShouldEqual, 11000)
		So(i.Store.PriSize, ShouldEqual, 13000)
	})
	Convey("Create index line from a complete index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 1 1234 3 11000 13000")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 1234)
		So(i.Docs.Deleted, ShouldEqual, 3)
		So(i.Store.Size, ShouldEqual, 11000)
		So(i.Store.PriSize, ShouldEqual, 13000)
	})
	Convey("Create index line from a bad docs index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 1 a 3 11000 13000")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 0)
		So(i.Docs.Deleted, ShouldEqual, 3)
		So(i.Store.Size, ShouldEqual, 11000)
		So(i.Store.PriSize, ShouldEqual, 13000)
	})
	Convey("Create index line from a bad deletes index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 1 1234 a 11000 13000")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 1234)
		So(i.Docs.Deleted, ShouldEqual, 0)
		So(i.Store.Size, ShouldEqual, 11000)
		So(i.Store.PriSize, ShouldEqual, 13000)
	})
	Convey("Create index line from a kinda short index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 1 1234")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 1234)
		So(i.Docs.Deleted, ShouldEqual, 0)
		So(i.Store.Size, ShouldEqual, 0)
		So(i.Store.PriSize, ShouldEqual, 0)
	})
	Convey("Create index line from a kinda sorta short index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 1 1234 3")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 1234)
		So(i.Docs.Deleted, ShouldEqual, 3)
		So(i.Store.Size, ShouldEqual, 0)
		So(i.Store.PriSize, ShouldEqual, 0)
	})
	Convey("Create index line from a short index listing", t, func() {
		i, err := NewCatIndexInfo("red foo-2000-01-01-bar 2 1")
		So(err, ShouldBeNil)
		So(i.Health, ShouldEqual, "red")
		So(i.Name, ShouldEqual, "foo-2000-01-01-bar")
		So(i.Shards, ShouldEqual, 2)
		So(i.Replicas, ShouldEqual, 1)
		So(i.Docs.Count, ShouldEqual, 0)
		So(i.Docs.Deleted, ShouldEqual, 0)
		So(i.Store.Size, ShouldEqual, 0)
		So(i.Store.PriSize, ShouldEqual, 0)
	})
}
