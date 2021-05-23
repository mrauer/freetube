package lib

import (
	"google.golang.org/api/youtube/v3"
)

// https://github.com/googleapis/google-api-go-client/blob/master/youtube/v3/youtube-gen.go

func SubscriptionsList(service *youtube.Service, part string, channelId string, hl string, maxResults int64, mine bool, onBehalfOfContentOwner string, pageToken string, playlistId string) *youtube.SubscriptionListResponse {
	call := service.Subscriptions.List([]string{part})
	if channelId != "" {
		call = call.ChannelId(channelId)
	}
	if mine != false {
		call = call.Mine(true)
	}
	call = call.MaxResults(maxResults)
	response, err := call.Do()
	HandleError(err, "")
	return response
}

func SearchList(service *youtube.Service, part string, channelId string, maxResults int64) *youtube.SearchListResponse {
	call := service.Search.List([]string{part})
	if channelId != "" {
		call = call.ChannelId(channelId)
	}

	call = call.Type("video")

	call = call.MaxResults(maxResults)
	response, err := call.Do()
	HandleError(err, "")
	return response
}
