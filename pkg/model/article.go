package model

type Article struct {
	*BaseModel
	Title       string   `bson:"title"`
	Slug        string   `bson:"slug"`
	Description string   `bson:"description"`
	Tags        []string `bson:"tags"`
	Content     string   `bson:"content"`
}
