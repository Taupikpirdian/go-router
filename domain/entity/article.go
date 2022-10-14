package entity

import (
	"time"
)

type Article struct {
	id          int
	slug        string
	lang        string
	category    string
	title       string
	text        string
	date        time.Time
	banner      string
	author      string
	thumbs      string
	isHighlight bool
}

type DTOArticle struct {
	Id          int
	Slug        string
	Lang        string
	Category    string
	Title       string
	Text        string
	Date        time.Time
	Banner      string
	Author      string
	Thumbs      string
	IsHighLight bool
}

// func create data
func NewCreateArticle(data DTOArticle) (*Article, error) {
	dataArticle := &Article{
		id:          data.Id,
		slug:        data.Slug,
		lang:        data.Lang,
		category:    data.Category,
		title:       data.Title,
		text:        data.Text,
		date:        data.Date,
		banner:      data.Banner,
		author:      data.Author,
		thumbs:      data.Thumbs,
		isHighlight: data.IsHighLight,
	}

	// err := dataArticle.validate()
	// if err != nil {
	// 	return nil, err
	// }

	return dataArticle, nil
}

func NewCreateArticleSingle(id int, slug string, lang string, category string, title string, text string, date time.Time, banner string, author string, thumbs string, is_high_light bool) (*Article, error) {
	dataArticle := &Article{
		id:          id,
		slug:        slug,
		lang:        lang,
		category:    category,
		title:       title,
		text:        text,
		date:        date,
		banner:      banner,
		author:      author,
		thumbs:      thumbs,
		isHighlight: is_high_light,
	}

	return dataArticle, nil
}

// func get data
func FetchDataArticleFromDB(dataDTO DTOArticle) *Article {
	dataArticle := &Article{
		id:          dataDTO.Id,
		slug:        dataDTO.Slug,
		lang:        dataDTO.Lang,
		category:    dataDTO.Category,
		title:       dataDTO.Title,
		text:        dataDTO.Text,
		date:        dataDTO.Date,
		banner:      dataDTO.Banner,
		author:      dataDTO.Author,
		thumbs:      dataDTO.Thumbs,
		isHighlight: dataDTO.IsHighLight,
	}

	return dataArticle
}

/*
	create func Getter
	supaya data struct utama bisa di akses dari luar
*/
func (data *Article) GetIdArtikel() int {
	return data.id
}

func (data *Article) GetSlugArtikel() string {
	return data.slug
}

func (data *Article) GetLangArtikel() string {
	return data.lang
}

func (data *Article) GetCategoryArtikel() string {
	return data.category
}

func (data *Article) GetTitleArtikel() string {
	return data.title
}

func (data *Article) GetTextArtikel() string {
	return data.text
}

func (data *Article) GetDateArtikel() time.Time {
	return data.date
}

func (data *Article) GetBannerArtikel() string {
	return data.banner
}

func (data *Article) GetAuthorArtikel() string {
	return data.author
}

func (data *Article) GetThumbsArtikel() string {
	return data.thumbs
}

func (data *Article) GetIsHighLightArtikel() bool {
	return data.isHighlight
}
