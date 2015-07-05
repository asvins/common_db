package commonDB

//for testing redis, make sure you have a local redis server up and running on port 6379

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestRedis(t *testing.T) { TestingT(t) }

//testing setup
type RedisSuite struct {
}

var _ = Suite(&RedisSuite{})

type MockStruct struct {
	Foo string
	Bar float64
}

func (s *RedisSuite) TestConnection(c *C) {
	db := NewRedisClient()
	stats, err := db.Ping().Result()

	c.Assert(err, IsNil)
	c.Assert(stats, NotNil)
}

func (s *RedisSuite) TestStore(c *C) {
	db := NewRedisClient()
	db.Set("foo", "bar", 0)
	result, _ := db.Get("foo").Result()
	c.Assert(result, Equals, "bar")
}

func (s *RedisSuite) TestStructStore(c *C) {
	db := NewRedisClient()
	var newObj MockStruct
	obj := MockStruct{"foo", 10.0}
	db.StoreStruct(obj.Foo, obj)
	db.GetStruct("foo", &newObj)
	c.Assert(newObj.Foo, Equals, obj.Foo)
	c.Assert(newObj.Bar, Equals, obj.Bar)
}
