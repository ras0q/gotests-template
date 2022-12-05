package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestNewFoo(t *testing.T) {
	t.Parallel()
	type args struct {
		repo RepoI
		srv  SrvI
	}
	tests := map[string]struct {
		args args
		want hoge
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := NewFoo(tt.args.repo, tt.args.srv)
			if diff := cmp.Diff(got, tt.want); len(diff) > 0 {
				t.Errorf("NewFoo() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_hoge_GetFromRepo(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo func(*gomock.Controller) RepoI
		srv  func(*gomock.Controller) SrvI
	}
	type args struct {
		id int
	}
	tests := map[string]struct {
		fields fields
		args   args
		want   Fuga
		err    error
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.fields
			ctrl := gomock.NewController(t)
			h := &hoge{
				repo: f.repo(ctrl),
				srv:  f.srv(ctrl),
			}
			got, err := h.GetFromRepo(tt.args.id)
			if !errors.Is(err, tt.err) {
				t.Errorf("hoge.GetFromRepo() error = %v, want %v", err, tt.err)
				return
			}
			if diff := cmp.Diff(got, tt.want); len(diff) > 0 {
				t.Errorf("hoge.GetFromRepo() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_hoge_DoWithSrv(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo func(*gomock.Controller) RepoI
		srv  func(*gomock.Controller) SrvI
	}
	tests := map[string]struct {
		fields fields
		err    error
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.fields
			ctrl := gomock.NewController(t)
			h := &hoge{
				repo: f.repo(ctrl),
				srv:  f.srv(ctrl),
			}
			if err := h.DoWithSrv(); !errors.Is(err, tt.err) {
				t.Errorf("hoge.DoWithSrv() error = %v, want %v", err, tt.err)
			}
		})
	}
}

func Test_hoge_Complecated(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo func(*gomock.Controller) RepoI
		srv  func(*gomock.Controller) SrvI
	}
	tests := map[string]struct {
		fields fields
		want   Fuga
		want1  bool
		err    error
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.fields
			ctrl := gomock.NewController(t)
			h := &hoge{
				repo: f.repo(ctrl),
				srv:  f.srv(ctrl),
			}
			got, got1, err := h.Complecated()
			if !errors.Is(err, tt.err) {
				t.Errorf("hoge.Complecated() error = %v, want %v", err, tt.err)
				return
			}
			if diff := cmp.Diff(got, tt.want); len(diff) > 0 {
				t.Errorf("hoge.Complecated() mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(got1, tt.want1); len(diff) > 0 {
				t.Errorf("hoge.Complecated() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_hoge_WriteToBuffer(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo func(*gomock.Controller) RepoI
		srv  func(*gomock.Controller) SrvI
	}
	tests := map[string]struct {
		fields fields
		wantW  string
		err    error
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		name := name
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.fields
			ctrl := gomock.NewController(t)
			h := &hoge{
				repo: f.repo(ctrl),
				srv:  f.srv(ctrl),
			}
			w := &bytes.Buffer{}
			if err := h.WriteToBuffer(w); !errors.Is(err, tt.err) {
				t.Errorf("hoge.WriteToBuffer() error = %v, want %v", err, tt.err)
				return
			}
			gotW := w.String()
			if diff := cmp.Diff(gotW, tt.wantW); len(diff) > 0 {
				t.Errorf("hoge.WriteToBuffer() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
