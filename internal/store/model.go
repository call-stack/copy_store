package store

type ContentMap struct {
	Hash string
	TTL  string
}

type DataStore struct {
	URL     string `bson:"url,omitempty"`
	Content string `bson:"content, omitempty"`
}
