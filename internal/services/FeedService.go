package services

import (
	"github.com/mmcdole/gofeed"
	"rss-feed/internal/models"
)

// FeedService provides methods to fetch RSS feeds.
type FeedService struct{}

// NewFeedService creates a new instance of FeedService.
func NewFeedService() *FeedService {
	return &FeedService{}
}

// FetchFeeds fetches RSS feeds from the given URL.
func (s *FeedService) FetchFeeds(url string) ([]models.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}

	var feeds []models.Feed
	for _, item := range feed.Items {
		feeds = append(feeds, models.Feed{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PublishedAt: *item.PublishedParsed,
			Read:        false,
		})
	}

	return feeds, nil
}
