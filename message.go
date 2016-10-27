package reddit

// Message is a private message between users.
type Message struct {
	Author           string      `json:"author"`
	Body             string      `json:"body"`
	BodyHTML         string      `json:"body_html"`
	Context          string      `json:"context"`
	Created          int         `json:"created"`
	CreatedUtc       int         `json:"created_utc"`
	Dest             string      `json:"dest"`
	Distinguished    string      `json:"distinguished"`
	FirstMessage     interface{} `json:"first_message"`
	FirstMessageName interface{} `json:"first_message_name"`
	ID               string      `json:"id"`
	Name             string      `json:"name"`
	New              bool        `json:"new"`
	ParentID         interface{} `json:"parent_id"`
	Replies          string      `json:"replies"`
	Subject          string      `json:"subject"`
	Subreddit        interface{} `json:"subreddit"`
	WasComment       bool        `json:"was_comment"`
}

// GetInboxMessages retrieves a list of messages in the user's inbox. Requires the 'privatemessages' OAuth scope.
func (c *Client) GetInboxMessages() ([]*Message, error) {
	return c.getMessages("inbox")
}

// GetUnreadMessages retrieves a list of unread messages in the user's inbox. Requires the 'privatemessages' OAuth scope.
func (c *Client) GetUnreadMessages() ([]*Message, error) {
	return c.getMessages("unread")
}

// GetSentMessages retrieves a list of messages sent by the user. Requires the 'privatemessages' OAuth scope.
func (c *Client) GetSentMessages() ([]*Message, error) {
	return c.getMessages("sent")
}

func (c *Client) getMessages(where string) ([]*Message, error) {
	return nil, nil
}
