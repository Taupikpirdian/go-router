package mysql_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"try/go-router/domain/entity"
	"try/go-router/domain/repository"
)

// interaksi dengan DB
type ArticleRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

// build structnya, yang mengacu ke connection dan kontrak interface di repository
func NewArticleRepositoryMysqlInteractor(connectionDatabse *sql.DB) repository.ArticleRepository {
	return &ArticleRepositoryMysqlInteractor{dbConn: connectionDatabse}
}

// implementasi dari interface kontrak dalam bentuk method receiver
func (repo *ArticleRepositoryMysqlInteractor) Store(ctx context.Context, dataArticle *entity.Article) error {
	// code here when u want store data to db
	return nil
}

func (repo *ArticleRepositoryMysqlInteractor) GetAllData(ctx context.Context, lang string, slug_category string) ([]*entity.Article, error) {
	var (
		errMysql error
	)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT article.id, article_lang.slug, article_lang.lang, categories_lang.name as category, title, text, date, banner, author, thumbs, is_highlight FROM article JOIN article_lang ON article_lang.base_id = article.id JOIN categories_lang ON categories_lang.base_id = article.category_id WHERE article_lang.lang = ? AND categories_lang.lang = ? AND categories_lang.slug = ?"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, lang, lang, slug_category)

	if errMysql != nil {
		fmt.Println(errMysql)
		return nil, errMysql
	}

	dataArticleCollection := make([]*entity.Article, 0)
	for rows.Next() {
		var (
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
		)

		err := rows.Scan(&id, &slug, &lang, &category, &title, &text, &date, &banner, &author, &thumbs, &isHighlight)
		if err != nil {
			return nil, err
		}
		dataArticle := entity.FetchDataArticleFromDB(entity.DTOArticle{
			Id:          id,
			Slug:        slug,
			Lang:        lang,
			Category:    category,
			Title:       title,
			Text:        text,
			Date:        date,
			Banner:      banner,
			Author:      author,
			Thumbs:      thumbs,
			IsHighLight: isHighlight,
		})

		dataArticleCollection = append(dataArticleCollection, dataArticle)
	}
	defer rows.Close()

	return dataArticleCollection, nil
}

func (repo *ArticleRepositoryMysqlInteractor) GetDataBySlug(ctx context.Context, slug string) (*entity.Article, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT id, slug, article_lang.lang, categories_lang.name as category, title, text, date, banner, author, thumbs, is_highlight FROM article JOIN article_lang ON article_lang.base_id = article.id JOIN categories_lang ON categories_lang.base_id = article.category_id WHERE article_lang.slug = ?"

	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, slug)
	if errMysql != nil {
		return nil, errMysql
	}

	if rows.Next() {
		var (
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
		)
		// first row
		err := rows.Scan(&id, &slug, &lang, &category, &title, &text, &date, &banner, &author, &thumbs, &isHighlight)
		if err != nil {
			return nil, err
		}
		// panic(id)

		dataArticle, err := entity.NewCreateArticleSingle(id, slug, lang, category, title, text, date, banner, author, thumbs, isHighlight)
		if err != nil {
			return nil, errors.New("gagal mapping data redis")
		}

		return dataArticle, nil

	} else {
		fmt.Println("... GAGAL Data Tidak Ditemukan, dengan slug artikel: " + slug)
		return nil, nil
	}
}
