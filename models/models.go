package models

import (
	"time"

	"github.com/google/uuid"
)

type Owner struct {
	Fullname    string `json:"fullname"`
	Password    string `json:"password"`
	Role        string `json:"role,omitempty"` // omitempty will omit this field if empty
	PhoneNumber string `json:"phone_numbeer"`
	Gmail       string `json:"gmail"`
	Telegram    string `json:"telegram"`
	GitHub      string `json:"github"`
	LinkedIn    string `json:"linked_in"`
	LeetCode    string `json:"leetcode"`
	AboutMe     string `json:"about_me,omitempty"`
}

type LoginOwn struct {
	PhoneNumber string `json:"phone_numbeer"`
	Gmail       string `json:"gmail"`
	Password    string `json:"password"`
}

type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateCategoryReq struct {
	Name string `json:"name"`
}

type GetCategoresListResp struct {
	Categories []*Category
	Count      int32
}

type Subcategory struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type Article struct {
	ArticleID     uuid.UUID  `json:"article_id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	CategoryID    uuid.UUID  `json:"category_id"`
	SubCategoryID uuid.UUID  `json:"sub_category_id"`
}

type Viewer struct {
	ViewerID uuid.UUID `json:"viewer_id"`
	Fullname string    `json:"fullname"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

type CheckViewer struct {
	Gmail string `json:"gmail"`
}

type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ArticleID uuid.UUID `json:"article_id"`
	ViewerID  uuid.UUID `json:"viewer_id"`
}

type GetList struct {
	Page  int32 `json:page`
	Limit int32 `json:limit`
}

type Common struct {
	TableName  string `json:table_name`
	ColumnName string `json:column_name`
	ExpValue   any    `json:exp_value`
}

type CheckExistsResp struct {
	IatsExists bool
	Status   string
}
