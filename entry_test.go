package wlog

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func TestWithFiled(t *testing.T) {
	log := &Entry{
		Data: make(TYFields, 0),
	}
	//log.WithField("test", "test")
	assert.Equal(t, 1, len(log.Data))
}

func TestEntry(t *testing.T) {
	// log := NewEntry(nil, "msg")
	// log.MustAppend().Flush()
}

func BenchmarkWithFiled(b *testing.B) {
	//log := &Entry{
	//	//Data: make(TYFileds, 0),
	//}
	for i := 0; i < b.N; i++ {
		//log.WithField("test", "test")
	}

}

func BenchmarkWithFileds(b *testing.B) {
	//log := &Entry{
	//	Data: make(TYFields, 0),
	//}
	//v := TYFields{
	//	"a": 1,
	//	"b": "b",
	//	"c": true,
	//}
	for i := 0; i < b.N; i++ {
		//log.TYFields(v)
	}
}

func BenchmarkInt64(b *testing.B) {
	log := &Entry{
		Data: make(TYFields, 0),
	}
	for i := 0; i < b.N; i++ {
		log.Int64("test", 109)
	}
}

//func BenchmarkInt(b *testing.B) {
//	log := New()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next() {
//			log.WithField("foo1", "bar1").
//				WithField("foo2", "bar2")
//		}
//	})
//}

func TestFuncMap(t *testing.T) {
	m := make(map[string]string, 0)
	m["a"] = "a"
	m["b"] = "b"
	m["c"] = "c"
	fmt.Println("xxxx", m)
}
