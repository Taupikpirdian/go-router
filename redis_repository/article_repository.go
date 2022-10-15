package redis_repository

import (
	"context"
	"fmt"
	"try/go-router/domain/entity"
	"try/go-router/redis_repository/mapper"

	"github.com/go-redis/redis/v8"
)

type RepoArticleRedis struct {
	Conn *redis.Client
}

func NewRepoArticleRedisInteractor(Conn *redis.Client) *RepoArticleRedis {
	return &RepoArticleRedis{Conn: Conn}
}

func (repo *RepoArticleRedis) GetAttributeArticleBySlug(ctx context.Context, kodeArticle string) (*entity.Article, error) {
	var (
		checkErr error
	)

	/*
		CLI Redis:
		HGET "XXXX-25524000" "data_article"
	*/

	data, checkErr := repo.Conn.HGet(ctx, kodeArticle, mapper.MapGetRedisField()).Result()
	if checkErr == redis.Nil {
		fmt.Println("Redis is Empty")
		return nil, nil
	}

	fmt.Println("... Yeah found single data in Redis")
	dataArticle, err := mapper.MapFromJsonStringToDomainArticle(data)
	if err != nil {
		return nil, err
	}

	return dataArticle, nil
}

func (repo *RepoArticleRedis) GetAllData(ctx context.Context) (*entity.Article, error) {
	/*
		func ini sementara blm digunakan, karena blm nemu cara get ALL data di redis by field
	*/
	var (
		checkErr error
	)

	data, checkErr := repo.Conn.HGetAll(ctx, mapper.MapGetRedisField()).Result()
	if len(data) == 0 || checkErr == redis.Nil {
		fmt.Println("Redis is Empty")
		return nil, nil
	}

	fmt.Println("... Yeah found data in Redis")
	// dataArticle, err := mapper.MapFromJsonStringToDomainArticle(data)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}

func (repo *RepoArticleRedis) StoreOrUpdateData(ctx context.Context, data *entity.Article) error {
	fmt.Println("=> => => Process Store Data To Redis")
	_, err := repo.Conn.HSet(ctx, data.GetSlugArtikel(), mapper.MapGetRedisField(), mapper.MapSetBukuToString(data)).Result()

	if err != nil {
		return err
	}

	return nil
}
