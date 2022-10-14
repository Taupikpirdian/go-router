package mapper

import (
	"encoding/json"
	"errors"
	"try/go-router/domain/entity"
	"try/go-router/mapper_response"
)

func MapGetRedisField() string {
	return "data_article"
}

func MapFromJsonStringToDomainArticle(data string) (*entity.Article, error) {
	attributeArticle := new(mapper_response.ArticleResponseJson)
	err := json.Unmarshal([]byte(data), attributeArticle)
	if err != nil {
		return nil, err
	}

	article, err := entity.NewCreateArticleSingle(
		attributeArticle.Id,
		attributeArticle.Slug,
		attributeArticle.Lang,
		attributeArticle.Category,
		attributeArticle.Title,
		attributeArticle.Text,
		attributeArticle.Date,
		attributeArticle.Banner,
		attributeArticle.Author,
		attributeArticle.Thumbs,
		attributeArticle.IsHighLight,
	)

	if err != nil {
		return nil, errors.New("gagal mapping data redis")
	}

	return article, nil
}

func MapSetBukuToString(data *entity.Article) string {
	attrBuku := &mapper_response.ArticleResponseJson{
		Id:          data.GetIdArtikel(),
		Slug:        data.GetSlugArtikel(),
		Lang:        data.GetLangArtikel(),
		Category:    data.GetCategoryArtikel(),
		Title:       data.GetTitleArtikel(),
		Text:        data.GetTextArtikel(),
		Date:        data.GetDateArtikel(),
		Banner:      data.GetBannerArtikel(),
		Author:      data.GetAuthorArtikel(),
		Thumbs:      data.GetThumbsArtikel(),
		IsHighLight: data.GetIsHighLightArtikel(),
	}

	attrJson, _ := json.Marshal(attrBuku)

	return string(attrJson)
}
