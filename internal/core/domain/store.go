package domain

type Store struct {
	URL     string `bson:"url,omitempty"`
	Hash    string `bson:"hash,omitempty"`
	Content string `bson:"content, omitempty"`
	TTL     string `bson:"ttl,omitempty"`
}
