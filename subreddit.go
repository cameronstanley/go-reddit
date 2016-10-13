package reddit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Subreddit struct {
	AccountsActive          int    `json:"accounts_active"`
	BannerImg               string `json:"banner_img"`
	BannerSize              []int  `json:"banner_size"`
	CollapseDeletedComments bool   `json:"collapse_deleted_comments"`
	CommentScoreHideMins    int    `json:"comment_score_hide_mins"`
	Created                 int    `json:"created"`
	CreatedUtc              int    `json:"created_utc"`
	Description             string `json:"description"`
	DescriptionHTML         string `json:"description_html"`
	DisplayName             string `json:"display_name"`
	HeaderImg               string `json:"header_img"`
	HeaderSize              []int  `json:"header_size"`
	HeaderTitle             string `json:"header_title"`
	HideAds                 bool   `json:"hide_ads"`
	IconImg                 string `json:"icon_img"`
	IconSize                []int  `json:"icon_size"`
	ID                      string `json:"id"`
	KeyColor                string `json:"key_color"`
	Lang                    string `json:"lang"`
	Name                    string `json:"name"`
	Over18                  bool   `json:"over18"`
	PublicDescription       string `json:"public_description"`
	PublicDescriptionHTML   string `json:"public_description_html"`
	PublicTraffic           bool   `json:"public_traffic"`
	Quarantine              bool   `json:"quarantine"`
	ShowMedia               bool   `json:"show_media"`
	ShowMediaPreview        bool   `json:"show_media_preview"`
	SubmissionType          string `json:"submission_type"`
	SubmitLinkLabel         string `json:"submit_link_label"`
	SubmitText              string `json:"submit_text"`
	SubmitTextHTML          string `json:"submit_text_html"`
	SubmitTextLabel         string `json:"submit_text_label"`
	SubredditType           string `json:"subreddit_type"`
	Subscribers             int    `json:"subscribers"`
	SuggestedCommentSort    string `json:"suggested_comment_sort"`
	Title                   string `json:"title"`
	URL                     string `json:"url"`
	UserIsBanned            bool   `json:"user_is_banned"`
	UserIsContributor       bool   `json:"user_is_contributor"`
	UserIsModerator         bool   `json:"user_is_moderator"`
	UserIsMuted             bool   `json:"user_is_muted"`
	UserIsSubscriber        bool   `json:"user_is_subscriber"`
	UserSrThemeEnabled      bool   `json:"user_sr_theme_enabled"`
	WikiEnabled             bool   `json:"wiki_enabled"`
}

func (c *Client) GetDefaultSubreddits() ([]*Subreddit, error) {
	url := fmt.Sprintf("%s/subreddits/default.json", baseUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Kind string `json:"kind"`
		Data struct {
			Modhash  string `json:"modhash"`
			Children []struct {
				Kind string    `json:"kind"`
				Data Subreddit `json:"data"`
			} `json:"children"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var subreddits []*Subreddit
	for _, subreddit := range result.Data.Children {
		subreddits = append(subreddits, &subreddit.Data)
	}

	return subreddits, nil
}
