package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinary(t *testing.T) {
	type args struct {
		val   int
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "Not found 1",
			args:    args{val: 99, input: make([]int, 0)},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Not found 2",
			args:    args{val: 99, input: []int{11}},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Not found 3",
			args:    args{val: 99, input: []int{4, 6}},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Not found 4",
			args:    args{val: 99, input: []int{1, 3, 4, 6, 11, 15, 20}},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Found 1",
			args:    args{val: 11, input: []int{11}},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "Found 2",
			args:    args{val: 6, input: []int{4, 6}},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Found 3",
			args:    args{val: 11, input: []int{1, 3, 4, 6, 11, 15, 20}},
			want:    4,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Binary(tt.args.val, tt.args.input)
			assert.Equal(t, got, tt.want)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}

func TestBinaryRecursive(t *testing.T) {
	type args struct {
		val   int
		input []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "Not found 1",
			args:    args{val: 99, input: make([]int, 0)},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Not found 2",
			args:    args{val: 99, input: []int{11}},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Not found 3",
			args:    args{val: 99, input: []int{4, 6}},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Not found 4",
			args:    args{val: 99, input: []int{1, 3, 4, 6, 11, 15, 20}},
			want:    0,
			wantErr: ErrSearchNotFound,
		},
		{
			name:    "Found 1",
			args:    args{val: 11, input: []int{11}},
			want:    0,
			wantErr: nil,
		},
		{
			name:    "Found 2",
			args:    args{val: 6, input: []int{4, 6}},
			want:    1,
			wantErr: nil,
		},
		{
			name:    "Found 3",
			args:    args{val: 11, input: []int{1, 3, 4, 6, 11, 15, 20}},
			want:    4,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BinaryRecursive(tt.args.val, tt.args.input)
			assert.Equal(t, got, tt.want)

			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}
