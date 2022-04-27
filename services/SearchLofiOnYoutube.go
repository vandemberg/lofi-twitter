package services

import (
	"twitter-lofi/entities"
	"twitter-lofi/infra"
)

type SearchLofiOnYoutube struct {
}

func (s *SearchLofiOnYoutube) Search(search string) []entities.Lofi {
	var lofiList []entities.Lofi

	youtubeList := infra.Youtube(search)

	for _, item := range youtubeList {
		if item.Id.Kind == "youtube#video" {
			lofiList = append(lofiList, entities.Lofi{
				Title:           item.Snippet.Title,
				VideoId:         item.Id.VideoId,
				SecondsDuration: 0,
			})
		}
	}

	return lofiList
}
