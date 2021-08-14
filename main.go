package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"os/exec"

	"github.com/gosimple/slug"
	"github.com/mrauer/freetube/lib"
	"google.golang.org/api/youtube/v3"
)

const (
	MAX_VIDEOS_DOWNLOAD       = 3
	MAX_PLAYLISTS_IN_RESPONSE = 100
)

var (
	method = flag.String("method", "list", "The API method to execute. (List is the only method that this sample currently supports.")

	channelId              = flag.String("channelId", "", "Retrieve playlists for this channel. Value is a YouTube channel ID.")
	hl                     = flag.String("hl", "", "Retrieve localized resource metadata for the specified application language.")
	maxResults             = flag.Int64("maxResults", MAX_PLAYLISTS_IN_RESPONSE, "The maximum number of playlist resources to include in the API response.")
	mine                   = flag.Bool("mine", true, "List playlists for authenticated user's channel. Default: false.")
	onBehalfOfContentOwner = flag.String("onBehalfOfContentOwner", "", "Indicates that the request's auth credentials identify a user authorized to act on behalf of the specified content owner.")
	pageToken              = flag.String("pageToken", "", "Token that identifies a specific page in the result set that should be returned.")
	part                   = flag.String("part", "snippet", "Comma-separated list of playlist resource parts that API response will include.")
	playlistId             = flag.String("playlistId", "", "Retrieve information about this playlist.")
)

func main() {
	flag.Parse()
	reader := bufio.NewReader(os.Stdin)

	client := lib.GetClient(youtube.YoutubeReadonlyScope)

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	subscriptions := lib.SubscriptionsList(service, "snippet", *channelId, *hl, *maxResults, *mine, *onBehalfOfContentOwner, *pageToken, *playlistId)

	choices := make(map[int]string)
	for idx, subscription := range subscriptions.Items {
		fmt.Println(fmt.Sprintf("%d : %s (%s)", idx, subscription.Snippet.Title, subscription.Snippet.ResourceId.ChannelId))
		choices[idx] = subscription.Snippet.ResourceId.ChannelId
	}

	fmt.Println("")
	fmt.Print("Choose from the above subscriptions which one to download all videos from: ")
	subscription_id_str, _ := reader.ReadString('\n')
	subscription_id, _ := strconv.Atoi(strings.TrimSuffix(subscription_id_str, "\n"))
	fmt.Println("")

	videos := lib.SearchList(service, "id,snippet", choices[subscription_id], MAX_VIDEOS_DOWNLOAD)

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	for _, video := range videos.Items {
		videoTitle, _, _ := transform.String(t, video.Snippet.Title)

		fmt.Println(fmt.Sprintf("Downloading %s", videoTitle))

		cmd := exec.Command("youtubedr", "download", "-o", fmt.Sprintf("%s.mp4", slug.Make(videoTitle)), "-d", "videos", "-q", "hd720", video.Id.VideoId)
		err := cmd.Run()

		if err != nil {
			fmt.Println(fmt.Printf("Error downloading video [%s]", err.Error()))
		}
	}
}
