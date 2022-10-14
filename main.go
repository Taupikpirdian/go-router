package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"try/go-router/domain/repository"
	"try/go-router/mapper_response"
	"try/go-router/mysql_repository"
	"try/go-router/pkg/mysql_connection"
	"try/go-router/pkg/redis_connection"
	"try/go-router/redis_repository"

	"github.com/gorilla/mux"
)

var (
	connectionDatabase     = mysql_connection.InitMysqlDB()
	articleRepositoryMysql = mysql_repository.NewArticleRepositoryMysqlInteractor(connectionDatabase)
	connectionRedis        = redis_connection.InitRedisClient()
	articleRepositoryRedis = redis_repository.NewRepoArticleRedisInteractor(connectionRedis)
	ctx                    = context.Background()
)

// for wadah kontrak interface
type ArticleLogicFactoryHandler struct {
	articleRepository repository.ArticleRepository
}

// initiate kontrak
func NewArticleLogicFactoryHandler(repoArticleImplementation repository.ArticleRepository) *ArticleLogicFactoryHandler {
	return &ArticleLogicFactoryHandler{articleRepository: repoArticleImplementation}
}

// for wadah kontrak interface
type ArticleLogicRedisFactoryHandler struct {
	articleRepository repository.ArticleRedisRepository
}

// initiate kontrak
func NewArticleLogicRedisFactoryHandler(repoArticleImplementation repository.ArticleRedisRepository) *ArticleLogicRedisFactoryHandler {
	return &ArticleLogicRedisFactoryHandler{articleRepository: repoArticleImplementation}
}

func main() {
	r := mux.NewRouter()
	/*
		untuk url /article tidak implement redis, karena apa dibutuhkan untuk data bentuk collection ?
		ini bagaimana supaya dinamis with param lang, en or id
	*/
	r.HandleFunc("/article", GetAllArticle("id", "makanan"))
	/*
		maaf blm tau cara assign param di url, dapetnya cara ini setelah browsing
		url /article/id sudah implement redis
	*/
	r.HandleFunc("/article/detail", GetArticleBySlug("the-bad-impact-of-instant-noodles"))

	http.ListenAndServe(":8080", r)
}

/*
function route
*/
func GetAllArticle(lang string, slug_category string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
		dataAllArticle, errGetData := handlerRepo.articleRepository.GetAllData(ctx, lang, slug_category)
		if errGetData != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		}

		responseAllArticle := make([]mapper_response.ArticleResponseJson, 0)
		for _, article := range dataAllArticle {
			articleResponse := mapper_response.BuildMapJSONFromDomain(article)
			responseAllArticle = append(responseAllArticle, articleResponse)
		}

		// return response
		response, _ := json.Marshal(responseAllArticle)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func GetArticleBySlug(slug string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get data from redis
		handlerRepoRedis := NewArticleLogicRedisFactoryHandler(articleRepositoryRedis)
		result, errGetDataRedis := handlerRepoRedis.articleRepository.GetAttributeArticleByKode(ctx, slug)
		if errGetDataRedis != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		}

		if result != nil {
			articleResponse := mapper_response.BuildMapJSONFromDomain(result)
			// return response
			response, _ := json.Marshal(articleResponse)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}

		// get data from db
		handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
		dataArticle, errGetData := handlerRepo.articleRepository.GetDataBySlug(ctx, slug)
		if errGetData != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "DATA DI DB TIDAK DITEMUKAN")
			return
		}
		articleResponse := mapper_response.BuildMapJSONFromDomain(dataArticle)

		// set to redis
		errStoreRepoRedis := handlerRepoRedis.articleRepository.StoreOrUpdateData(ctx, dataArticle)
		if errStoreRepoRedis != nil {
			fmt.Fprintf(w, "GAGAL CREATE ARTICLE TO REDIS ADA KESALAHAN DALAM PENYIMPANAN KE REDIS")
			return
			// panic(errStoreRepoRedis)
		}

		// return response
		response, _ := json.Marshal(articleResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
