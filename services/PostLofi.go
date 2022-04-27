package services

import (
	"twitter-lofi/entities"
	"twitter-lofi/infra"
)

type PostLofi struct {
}

func (p *PostLofi) Post(lofi entities.Lofi) {
	youtubeUrl := "https://www.youtube.com/watch?v="
	infra.Twitter(youtubeUrl + lofi.VideoId)
}
