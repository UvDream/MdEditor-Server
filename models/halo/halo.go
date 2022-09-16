package halo

type UserHalo struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Url      string `form:"url"`
}

type ArticleHaloResponse struct {
	Url   string `json:"url"`
	Token string `json:"token"`
	ArticleHalo
}
type ArticleHalo struct {
	ID              int64  `json:"id"`
	Title           string `json:"title"`
	EditorType      string `json:"editorType"` //MARKDOWN RICHTEXT
	Slug            string `json:"slug"`       //别名
	Status          string `json:"status"`     //PUBLISHED DRAFT INTIMATE RECYCLE
	CategoryIds     []int  `json:"categoryIds"`
	TagIds          []int  `json:"tagIds"`
	Content         string `json:"content"`         //html
	OriginalContent string `json:"originalContent"` //markdown
	FormatContent   string `json:"formatContent"`   //html
	Topped          bool   `json:"topped"`          //是否置顶
}
