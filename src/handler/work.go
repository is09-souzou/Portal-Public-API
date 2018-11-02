package handler

// Work response work struct
type Work struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Tags        *[]string `json:"tags"`
	ImageURL    *string   `json:"imageUrl"`
	Description string    `json:"description"`
	CreatedAt   int       `json:"createdAt"`
}

// WorkConnection response work connection struct
type WorkConnection struct {
	Items             []Work  `json:"items"`
	ExclusiveStartKey *string `json:"exclusiveStartKey"`
}
