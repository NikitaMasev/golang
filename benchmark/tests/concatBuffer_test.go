package tests

import (
	"testing"
	"bytes"
    "fmt"
	"github.com/NikitaMasev/golang/benchmark/data"
)

func BenchmarkConcatBuffer(b *testing.B)  {
    b.ReportAllocs()
    for i := 0;i < b.N;i++ {
       key:=bytes.Buffer{}
       key.WriteString(data.RecordData.A+"_")
       key.WriteString(data.RecordData.B+"_")
       key.WriteString(data.RecordData.C+"_")
       key.WriteString(data.RecordData.D+"_")
       key.WriteString(data.RecordData.E+"_")
       key.WriteString(data.RecordData.F)
       fmt.Sprintf("r_%v",key)
    }
    
}

/*
       BenchmarkConcatBuffer-2   	  100000	     11549 ns/op	    1008 B/op	       4 allocs/op
*/
