package main

import (
	"bytes"
	"testing"

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
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if diff := cmp.Diff(tt.want, NewFoo(tt.args.repo, tt.args.srv)); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func Test_hoge_GetFromRepo(t *testing.T) {
	t.Parallel()
	type args struct {
		id int
	}
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	type setupFieldsFunc func(t *testing.T, args args, want Fuga) fields
	tests := map[string]struct {
		args        args
		want        Fuga
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t, tt.args, tt.want)
			h := &hoge{
				repo: f.repo,
				srv:  f.srv,
			}
			got, err := h.GetFromRepo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func Test_hoge_DoWithSrv(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	type setupFieldsFunc func(t *testing.T) fields
	tests := map[string]struct {
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t)
			h := &hoge{
				repo: f.repo,
				srv:  f.srv,
			}
			if err := h.DoWithSrv(); (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
		})
	}
}

func Test_hoge_Complecated(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	type setupFieldsFunc func(t *testing.T, want Fuga, want1 bool) fields
	tests := map[string]struct {
		want        Fuga
		want1       bool
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t, tt.want, tt.want1)
			h := &hoge{
				repo: f.repo,
				srv:  f.srv,
			}
			got, got1, err := h.Complecated()
			if (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
			if diff := cmp.Diff(tt.want1, got1); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}

func Test_hoge_WriteToBuffer(t *testing.T) {
	t.Parallel()
	type fields struct {
		repo RepoI
		srv  SrvI
	}
	type setupFieldsFunc func(t *testing.T, wantW string) fields
	tests := map[string]struct {
		wantW       string
		setupFields setupFieldsFunc
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			f := tt.setupFields(t, tt.wantW)
			h := &hoge{
				repo: f.repo,
				srv:  f.srv,
			}
			w := &bytes.Buffer{}
			if err := h.WriteToBuffer(w); (err != nil) != tt.wantErr {
				t.Errorf("wantErr is %t, but err is %v", tt.wantErr, err)
			}
			if diff := cmp.Diff(tt.wantW,
				w.String()); len(diff) > 0 {
				t.Errorf("Compare value is mismatch (-want +got):%s\n", diff)
			}
		})
	}
}
