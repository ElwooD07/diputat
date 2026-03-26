package models

// SocialContact stores public social profile references.
type SocialContact struct {
	Platform string `json:"platform"`
	URL      string `json:"url"`
	Handle   string `json:"handle,omitempty"`
}

// OfficialContact stores direct official contact channels.
type OfficialContact struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// ContactSet groups social and official contacts for a public official.
type ContactSet struct {
	Social   []SocialContact   `json:"social,omitempty"`
	Official []OfficialContact `json:"official,omitempty"`
}

// Official represents a person whose public claims are tracked.
type Official struct {
	ID       string     `json:"_id"`
	Name     string     `json:"name"`
	Handle   string     `json:"handle,omitempty"`
	Position string     `json:"position"`
	Party    string     `json:"party,omitempty"`
	Contacts ContactSet `json:"contacts,omitempty"`
	Bio      string     `json:"bio,omitempty"`
	PhotoURL string     `json:"photo_url,omitempty"`
	Flag     bool       `json:"flag,omitempty"`
	FlagNote string     `json:"flag_reason,omitempty"`
	Metadata Metadata   `json:"metadata,omitempty"`
}
