package mapper_response

import (
	"time"
	"try/go-router/domain/entity"
)

type ArticleResponseJson struct {
	Id          int       `json:"id"`
	Slug        string    `json:"slug"`
	Lang        string    `json:"lang"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Text        string    `json:"text"`
	Date        time.Time `json:"date"`
	Banner      string    `json:"banner"`
	Author      string    `json:"author"`
	Thumbs      string    `json:"thumbs"`
	IsHighLight bool      `json:"is_highlifgt"`
}

func BuildMapJSONFromDomain(domain *entity.Article) ArticleResponseJson {
	return ArticleResponseJson{
		Id:          domain.GetIdArtikel(),
		Slug:        domain.GetSlugArtikel(),
		Lang:        domain.GetLangArtikel(),
		Category:    domain.GetCategoryArtikel(),
		Title:       domain.GetTitleArtikel(),
		Text:        domain.GetTextArtikel(),
		Date:        domain.GetDateArtikel(),
		Banner:      domain.GetBannerArtikel(),
		Author:      domain.GetAuthorArtikel(),
		Thumbs:      domain.GetThumbsArtikel(),
		IsHighLight: domain.GetIsHighLightArtikel(),
	}
}
