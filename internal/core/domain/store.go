package domain

type Store struct {
	Hash    string `bson:"hash,omitempty"`
	Content string `bson:"content, omitempty"`
	TTL     string `bson:"ttl,omitempty"`
}
