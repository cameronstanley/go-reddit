package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	ContestMode      bool          `json:"contest_mode"`
	BannedBy         interface{}   `json:"banned_by"`
	Domain           string        `json:"domain"`
	Subreddit        string        `json:"subreddit"`
	SelftextHTML     interface{}   `json:"selftext_html"`
	Selftext         string        `json:"selftext"`
	Likes            interface{}   `json:"likes"`
	SuggestedSort    interface{}   `json:"suggested_sort"`
	UserReports      []interface{} `json:"user_reports"`
	SecureMedia      interface{}   `json:"secure_media"`
	Saved            bool          `json:"saved"`
	ID               string        `json:"id"`
	Gilded           int           `json:"gilded"`
	SecureMediaEmbed struct {
	} `json:"secure_media_embed"`
	Clicked             bool          `json:"clicked"`
	ReportReasons       interface{}   `json:"report_reasons"`
	Author              string        `json:"author"`
	Media               interface{}   `json:"media"`
	Name                string        `json:"name"`
	Score               int           `json:"score"`
	ApprovedBy          interface{}   `json:"approved_by"`
	Over18              bool          `json:"over_18"`
	RemovalReason       interface{}   `json:"removal_reason"`
	Hidden              bool          `json:"hidden"`
	Thumbnail           string        `json:"thumbnail"`
	SubredditID         string        `json:"subreddit_id"`
	Edited              bool          `json:"edited"`
	LinkFlairCSSClass   interface{}   `json:"link_flair_css_class"`
	AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
	Downs               int           `json:"downs"`
	ModReports          []interface{} `json:"mod_reports"`
	Archived            bool          `json:"archived"`
	MediaEmbed          struct {
	} `json:"media_embed"`
	IsSelf          bool        `json:"is_self"`
	HideScore       bool        `json:"hide_score"`
	Permalink       string      `json:"permalink"`
	Locked          bool        `json:"locked"`
	Stickied        bool        `json:"stickied"`
	Created         int         `json:"created"`
	URL             string      `json:"url"`
	AuthorFlairText interface{} `json:"author_flair_text"`
	Quarantine      bool        `json:"quarantine"`
	Title           string      `json:"title"`
	CreatedUtc      int         `json:"created_utc"`
	LinkFlairText   interface{} `json:"link_flair_text"`
	Distinguished   interface{} `json:"distinguished"`
	NumComments     int         `json:"num_comments"`
	Visited         bool        `json:"visited"`
	NumReports      interface{} `json:"num_reports"`
	Ups             int         `json:"ups"`
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
