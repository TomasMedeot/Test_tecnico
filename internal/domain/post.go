package domain

// Post represents a post in the system with basic attributes like ID, UserID, Likes, and Content.
type Post struct {
	ID      interface{} `json:"id"`      // Unique identifier for the post
	UserID  int         `json:"userid"`  // Identifier of the user who created the post
	Likes   int         `json:"likes"`   // Number of likes the post has received
	Content string      `json:"content"` // Text content of the post
}

// NewPost creates a new Post instance with the provided ID, UserID, and Content.
func NewPost(ID interface{}, UserID int, Likes int, Content string) *Post {
	if ID == nil {
		ID = 0
	}
	if Likes < 0 {
		Likes = 0
	}
	return &Post{
		ID:      ID,
		UserID:  UserID,
		Likes:   Likes,
		Content: Content,
	}
}
