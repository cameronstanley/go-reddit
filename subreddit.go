package reddit

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
)

type Subreddits struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash string `json:"modhash"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				BannerImg string `json:"banner_img"`
				UserSrThemeEnabled bool `json:"user_sr_theme_enabled"`
				SubmitTextHTML string `json:"submit_text_html"`
				UserIsBanned interface{} `json:"user_is_banned"`
				WikiEnabled bool `json:"wiki_enabled"`
				ShowMedia bool `json:"show_media"`
				ID string `json:"id"`
				SubmitText string `json:"submit_text"`
				DisplayName string `json:"display_name"`
				HeaderImg string `json:"header_img"`
				DescriptionHTML string `json:"description_html"`
				Title string `json:"title"`
				CollapseDeletedComments bool `json:"collapse_deleted_comments"`
				Over18 bool `json:"over18"`
				PublicDescriptionHTML string `json:"public_description_html"`
				IconSize []int `json:"icon_size"`
				SuggestedCommentSort interface{} `json:"suggested_comment_sort"`
				IconImg string `json:"icon_img"`
				HeaderTitle interface{} `json:"header_title"`
				Description string `json:"description"`
				UserIsMuted interface{} `json:"user_is_muted"`
				SubmitLinkLabel interface{} `json:"submit_link_label"`
				AccountsActive interface{} `json:"accounts_active"`
				PublicTraffic bool `json:"public_traffic"`
				HeaderSize []int `json:"header_size"`
				Subscribers int `json:"subscribers"`
				SubmitTextLabel interface{} `json:"submit_text_label"`
				Lang string `json:"lang"`
				UserIsModerator interface{} `json:"user_is_moderator"`
				KeyColor string `json:"key_color"`
				Name string `json:"name"`
				Created int `json:"created"`
				URL string `json:"url"`
				Quarantine bool `json:"quarantine"`
				HideAds bool `json:"hide_ads"`
				CreatedUtc int `json:"created_utc"`
				BannerSize interface{} `json:"banner_size"`
				UserIsContributor interface{} `json:"user_is_contributor"`
				PublicDescription string `json:"public_description"`
				ShowMediaPreview bool `json:"show_media_preview"`
				CommentScoreHideMins int `json:"comment_score_hide_mins"`
				SubredditType string `json:"subreddit_type"`
				SubmissionType string `json:"submission_type"`
				UserIsSubscriber interface{} `json:"user_is_subscriber"`
			} `json:"data"`
		} `json:"children"`
		After string `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

func (c *Client) GetDefaultSubreddits() {
  url :=  fmt.Sprintf("%s/subreddits/default.json", baseUrl)
  resp, _ := http.Get(url)
  defer resp.Body.Close()

  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Print(string(body))

  var subreddits Subreddits
  json.NewDecoder(resp.Body).Decode(&subreddits)
}
