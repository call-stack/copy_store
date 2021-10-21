package domain

type Store struct {
	URL     string `bson:"url,omitempty"`
	Content string `bson:"content, omitempty"`
	TTL     string `bson:"ttl,omitempty"`
}
