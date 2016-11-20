package tests

import (
	"testing"
	"strings"
    "fmt"
	"github.com/NikitaMasev/golang/benchmark/data"
)

func BenchmarkConcatJoin(b *testing.B)  {
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        var  s[] string
        strings.Join(s,data.RecordData.A+"_")
        strings.Join(s,data.RecordData.B+"_")
        strings.Join(s,data.RecordData.C+"_")
        strings.Join(s,data.RecordData.D+"_")
        strings.Join(s,data.RecordData.E+"_")
        strings.Join(s,data.RecordData.F)
        fmt.Sprintf("r_%v",s)
    }    
}

/*
    BenchmarkConcatJoin-2     	 3000000	       488 ns/op	      36 B/op	       2 allocs/op
    
*/
