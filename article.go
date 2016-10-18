package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	ApprovedBy          string        `json:"approved_by"`
	Archived            bool          `json:"archived"`
	Author              string        `json:"author"`
	AuthorFlairCSSClass string        `json:"author_flair_css_class"`
	AuthorFlairText     string        `json:"author_flair_text"`
	BannedBy            string        `json:"banned_by"`
	Clicked             bool          `json:"clicked"`
	ContestMode         bool          `json:"contest_mode"`
	Created             int           `json:"created"`
	CreatedUtc          int           `json:"created_utc"`
	Distinguished       string        `json:"distinguished"`
	Domain              string        `json:"domain"`
	Downs               int           `json:"downs"`
	Edited              bool          `json:"edited"`
	Gilded              int           `json:"gilded"`
	Hidden              bool          `json:"hidden"`
	HideScore           bool          `json:"hide_score"`
	ID                  string        `json:"id"`
	IsSelf              bool          `json:"is_self"`
	Likes               bool          `json:"likes"`
	LinkFlairCSSClass   string        `json:"link_flair_css_class"`
	LinkFlairText       string        `json:"link_flair_text"`
	Locked              bool          `json:"locked"`
	Media               Media         `json:"media"`
	MediaEmbed          interface{}   `json:"media_embed"`
	ModReports          []interface{} `json:"mod_reports"`
	Name                string        `json:"name"`
	NumComments         int           `json:"num_comments"`
	NumReports          int           `json:"num_reports"`
	Over18              bool          `json:"over_18"`
	Permalink           string        `json:"permalink"`
	Quarantine          bool          `json:"quarantine"`
	RemovalReason       interface{}   `json:"removal_reason"`
	ReportReasons       []interface{} `json:"report_reasons"`
	Saved               bool          `json:"saved"`
	Score               int           `json:"score"`
	SecureMedia         interface{}   `json:"secure_media"`
	SecureMediaEmbed    interface{}   `json:"secure_media_embed"`
	SelftextHTML        string        `json:"selftext_html"`
	Selftext            string        `json:"selftext"`
	Stickied            bool          `json:"stickied"`
	Subreddit           string        `json:"subreddit"`
	SubredditID         string        `json:"subreddit_id"`
	SuggestedSort       string        `json:"suggested_sort"`
	Thumbnail           string        `json:"thumbnail"`
	Title               string        `json:"title"`
	URL                 string        `json:"url"`
	Ups                 int           `json:"ups"`
	UserReports         []interface{} `json:"user_reports"`
	Visited             bool          `json:"visited"`
}

type articleListing struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Children []struct {
			Kind string  `json:"kind"`
			Data Article `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func (c *Client) GetHotArticles(subreddit string) ([]*Article, error) {
	return c.getArticles(subreddit, "hot")
}

func (c *Client) GetNewArticles(subreddit string) ([]*Article, error) {
	return c.getArticles(subreddit, "new")
}

func (c *Client) GetTopArticles(subreddit string) ([]*Article, error) {
	return c.getArticles(subreddit, "top")
}

func (c *Client) getArticles(subreddit string, sort string) ([]*Article, error) {
	url := fmt.Sprintf("%s/r/%s/%s.json", baseUrl, subreddit, sort)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result articleListing
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var articles []*Article
	for _, article := range result.Data.Children {
		articles = append(articles, &article.Data)
	}

	return articles, nil
}
