package http

import (
	"testing"

	"github.com/ArtisanCloud/MediaXCore/pkg/http/contract"
)

func Test_NewClient(t *testing.T) {
	helper, err := NewHttpClient(&contract.ClientConfig{})
	if err != nil {
		t.Error(err)
	}

	if helper == nil {
		t.Error(err)
	}

}
