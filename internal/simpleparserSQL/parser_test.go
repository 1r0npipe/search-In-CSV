package simpleparserSQL

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

func Test_isDigit(t *testing.T) {
	tests := []struct {
		ch rune
		want bool
	}{
		{ch: '0', want: true},
		{ch: '5', want: true},
		{ch: '9', want: true},
		{ch: 'f', want: false},
	}
	for _, tt := range tests {
		got := isDigit(tt.ch);
		assert.Equal(t, got, tt.want, "Not equil, error for %v !", string(tt.ch))
	}
}

func Test_isLetter(t *testing.T) {
	tests := []struct {
		ch rune
		want bool
	}{
		{ch: 'a', want: true},
		{ch: '5', want: false},
		{ch: 'D', want: true},
		{ch: '0', want: false},
	}
	for _, tt := range tests {
		got := isLetter(tt.ch);
		assert.Equal(t, got, tt.want, "Not equil, error for %v !", string(tt.ch))
	}
}