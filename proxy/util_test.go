package proxy

import (
	"net/url"
	"testing"

	"github.com/jhump/protoreflect/desc"
)

func parseURL(t *testing.T, rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		t.Fatal(err.Error())
	}
	return u
}

func newFileDescriptor(t *testing.T, file string) *desc.FileDescriptor {
	t.Helper()
	desc, err := desc.LoadFileDescriptor(file)
	if err != nil {
		t.Fatal(err.Error())
	}
	return desc
}