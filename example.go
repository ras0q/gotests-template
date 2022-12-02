package main

import "io"

type hoge struct {
	repo RepoI
	srv  SrvI
}

type Fuga struct {
	// id   int
	// name string
}

type RepoI interface {
	Get(id int) (Fuga, error)
}

type SrvI interface {
	Do() error
}

func NewFoo(repo RepoI, srv SrvI) hoge {
	return hoge{repo, srv}
}

func (h *hoge) GetFromRepo(id int) (Fuga, error) {
	return h.repo.Get(id)
}

func (h *hoge) DoWithSrv() error {
	return h.srv.Do()
}

func (h *hoge) Complecated() (Fuga, bool, error) {
	if err := h.srv.Do(); err != nil {
		return Fuga{}, false, err
	}

	id := 1

	f, err := h.repo.Get(id)
	if err != nil {
		return Fuga{}, false, err
	}

	return f, true, nil
}

func (h *hoge) WriteToBuffer(w io.Writer) error {
	if _, err := w.Write([]byte("hoge")); err != nil {
		return err
	}

	return nil
}
