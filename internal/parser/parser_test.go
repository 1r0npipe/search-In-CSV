package parser

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_isWhiteSpace(t *testing.T) {
	tests := []struct {
		ch rune
		want bool
	}{
		{ch: ' ', want: true},
		{ch: '\t', want: true},
		{ch: '\n', want: true},
		{ch: 'f', want: false},
	}
	for _, tt := range tests {
		got := isWhiteSpace(tt.ch);
		assert.Equal(t, got, tt.want, "Not equil, error for %v !", string(tt.ch))
	}
}
