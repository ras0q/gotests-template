package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFoo(t *testing.T) {
	t.Parallel()
	type args struct {
		repo RepoI
		srv  SrvI
	}
	tests := []struct {
		name  string
		args  args
		want  hoge
		setup func(args args, want hoge)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, NewFoo(tt.args.repo, tt.args.srv))
		})
	}
}

func Test_hoge_GetFromRepo(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	type args struct {
		id int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      Fuga
		setup     func(f fields, args args, want Fuga)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := &hoge{
				repo: tt.fields.repo,
				srv:  tt.fields.srv,
			}
			if tt.setup != nil {
				tt.setup(tt.fields, tt.args, tt.want)
			}
			got, err := h.GetFromRepo(tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_hoge_DoWithSrv(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	tests := []struct {
		name      string
		fields    fields
		setup     func(f fields)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := &hoge{
				repo: tt.fields.repo,
				srv:  tt.fields.srv,
			}
			if tt.setup != nil {
				tt.setup(tt.fields)
			}
			tt.assertion(t, h.DoWithSrv())
		})
	}
}

func Test_hoge_Complecated(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	tests := []struct {
		name      string
		fields    fields
		want      Fuga
		want1     bool
		setup     func(f fields, want Fuga, want1 bool)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := &hoge{
				repo: tt.fields.repo,
				srv:  tt.fields.srv,
			}
			if tt.setup != nil {
				tt.setup(tt.fields, tt.want, tt.want1)
			}
			got, got1, err := h.Complecated()
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func Test_hoge_WriteToBuffer(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	tests := []struct {
		name      string
		fields    fields
		wantW     string
		setup     func(f fields, wantW string)
		assertion assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			h := &hoge{
				repo: tt.fields.repo,
				srv:  tt.fields.srv,
			}
			if tt.setup != nil {
				tt.setup(tt.fields, tt.wantW)
			}
			w := &bytes.Buffer{}
			tt.assertion(t, h.WriteToBuffer(w))
			assert.Equal(t, tt.wantW, w.String())
		})
	}
}
