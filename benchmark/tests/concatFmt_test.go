package tests

import (
	"testing"
	"fmt"
	"github.com/NikitaMasev/golang/benchmark/data"
)

func BenchmarkConcatFmt(b *testing.B)  {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        key:=fmt.Sprintf("%v%v%v%v%v%v",
            data.RecordData.A+"_",
            data.RecordData.B+"_",
            data.RecordData.C+"_",
            data.RecordData.D+"_",
            data.RecordData.E+"_",
            data.RecordData.F)
            fmt.Sprintf("r_%v",key)
    }
}

/*
    BenchmarkConcatFmt-2      	 1000000	      1599 ns/op	     384 B/op	      14 allocs/op
*/
