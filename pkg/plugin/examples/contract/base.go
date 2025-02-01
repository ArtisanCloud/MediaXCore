package contract

import "github.com/ArtisanCloud/MediaXCore/pkg/plugin/core/contract"

type PublishRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

type PublishResponse struct {
	contract.BaseResponse
}
