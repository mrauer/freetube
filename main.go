package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mrauer/freetube/lib"
	"google.golang.org/api/youtube/v3"
)

var (
	method = flag.String("method", "list", "The API method to execute. (List is the only method that this sample currently supports.")

	channelId              = flag.String("channelId", "", "Retrieve playlists for this channel. Value is a YouTube channel ID.")
	hl                     = flag.String("hl", "", "Retrieve localized resource metadata for the specified application language.")
	maxResults             = flag.Int64("maxResults", 7, "The maximum number of playlist resources to include in the API response.")
	mine                   = flag.Bool("mine", false, "List playlists for authenticated user's channel. Default: false.")
	onBehalfOfContentOwner = flag.String("onBehalfOfContentOwner", "", "Indicates that the request's auth credentials identify a user authorized to act on behalf of the specified content owner.")
	pageToken              = flag.String("pageToken", "", "Token that identifies a specific page in the result set that should be returned.")
	part                   = flag.String("part", "snippet", "Comma-separated list of playlist resource parts that API response will include.")
	playlistId             = flag.String("playlistId", "", "Retrieve information about this playlist.")
)

func subscriptionsList(service *youtube.Service, part string, channelId string, hl string, maxResults int64, mine bool, onBehalfOfContentOwner string, pageToken string, playlistId string) *youtube.SubscriptionListResponse {
	call := service.Subscriptions.List([]string{part})
	if channelId != "" {
		call = call.ChannelId(channelId)
	}
	if mine != false {
		call = call.Mine(true)
	}
	call = call.MaxResults(maxResults)
	response, err := call.Do()
	lib.HandleError(err, "")
	return response
}

func main() {
	flag.Parse()

	if *channelId == "" && *mine == false && *playlistId == "" {
		log.Fatalf("You must either set a value for the channelId or playlistId flag or set the mine flag to 'true'.")
	}
	client := lib.GetClient(youtube.YoutubeReadonlyScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	subscriptions := subscriptionsList(service, "snippet,contentDetails", *channelId, *hl, *maxResults, *mine, *onBehalfOfContentOwner, *pageToken, *playlistId)

	for _, subscription := range subscriptions.Items {
		fmt.Println(fmt.Sprintf("%s - %s", subscription.Id, subscription.Snippet.Title))
	}
}
