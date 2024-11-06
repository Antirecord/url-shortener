package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateHash(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		val  string
	}{
		{
			name: "first",
			args: args{data: "testurl"},
			val:  "uuid-gen",
		},
		{
			name: "second",
			args: args{data: "test123"},
			val:  "uuid-gen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateHash(tt.args.data)
			assert.NotEmpty(t, got)
			assert.Len(t, got, 8)
			assert.NotEqual(t, got, tt.val)

		})
	}
}
