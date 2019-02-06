package main_test

import (
	"testing"

	s3 "github.com/som-poddar/aws-golang-benchmark"
)

func BenchmarkListBuckets(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s3.ListBuckets()
	}
}
