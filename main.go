package main

import (
	"flag"
	"fmt"
	"log"

	"os/exec"

	"github.com/mrauer/freetube/lib"
	"google.golang.org/api/youtube/v3"
)

const (
	SEARCH_LIST_MAX_RESULTS   = 8
	MAX_PLAYLISTS_IN_RESPONSE = 7
)

var (
	method = flag.String("method", "list", "The API method to execute. (List is the only method that this sample currently supports.")

	channelId              = flag.String("channelId", "", "Retrieve playlists for this channel. Value is a YouTube channel ID.")
	hl                     = flag.String("hl", "", "Retrieve localized resource metadata for the specified application language.")
	maxResults             = flag.Int64("maxResults", MAX_PLAYLISTS_IN_RESPONSE, "The maximum number of playlist resources to include in the API response.")
	mine                   = flag.Bool("mine", false, "List playlists for authenticated user's channel. Default: false.")
	onBehalfOfContentOwner = flag.String("onBehalfOfContentOwner", "", "Indicates that the request's auth credentials identify a user authorized to act on behalf of the specified content owner.")
	pageToken              = flag.String("pageToken", "", "Token that identifies a specific page in the result set that should be returned.")
	part                   = flag.String("part", "snippet", "Comma-separated list of playlist resource parts that API response will include.")
	playlistId             = flag.String("playlistId", "", "Retrieve information about this playlist.")
)

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

	subscriptions := lib.SubscriptionsList(service, "snippet", *channelId, *hl, *maxResults, *mine, *onBehalfOfContentOwner, *pageToken, *playlistId)

	for _, subscription := range subscriptions.Items {
		fmt.Println(fmt.Sprintf("%s - %s", subscription.Snippet.ResourceId.ChannelId, subscription.Snippet.Title))
	}

	videos := lib.SearchList(service, "id,snippet", "UCynFUJ4zUVuh3GX7bABTjGQ", SEARCH_LIST_MAX_RESULTS)

	for _, video := range videos.Items {
		fmt.Println(video.Id.VideoId)

		cmd := exec.Command("youtubedr", "download", "-d", "videos", "-q", "hd720", video.Id.VideoId)
		err := cmd.Run()

		if err != nil {
			fmt.Println(fmt.Printf("Error downloading video [%s]", err.Error()))
		}

		break
	}
}
