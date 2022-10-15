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
	/*
		- Cara penggunaan ada pada README.md
		- url /article/{lang} tidak implement ke redis, karena apa dibutuhkan untuk data bentuk collection ?
		- url /article/detail/{slug} sudah implement redis
		- untuk poin 3 di readme, saya tidak mencantumkan langnya, karena dengan slug saja itu sudah mewakili apakah itu bahasa indonesia atau bahasa inggris
	*/
	r := mux.NewRouter()
	r.HandleFunc("/article/{lang}", GetAllArticle)
	r.HandleFunc("/article/detail/{slug}", GetArticleBySlug)

	http.ListenAndServe(":8080", r)
}

/*
function route
*/
func GetAllArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	lang := params["lang"]
	category := r.URL.Query().Get("category")

	handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
	dataAllArticle, errGetData := handlerRepo.articleRepository.GetAllData(ctx, lang, category)
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

func GetArticleBySlug(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	slug := params["slug"]

	// get data from redis
	handlerRepoRedis := NewArticleLogicRedisFactoryHandler(articleRepositoryRedis)
	result, errGetDataRedis := handlerRepoRedis.articleRepository.GetAttributeArticleBySlug(ctx, slug)
	if errGetDataRedis != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
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
		return
	}
	if dataArticle == nil {
		fmt.Fprintf(w, "DATA DI DB TIDAK DITEMUKAN")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	articleResponse := mapper_response.BuildMapJSONFromDomain(dataArticle)

	// set to redis
	errStoreRepoRedis := handlerRepoRedis.articleRepository.StoreOrUpdateData(ctx, dataArticle)
	if errStoreRepoRedis != nil {
		fmt.Fprintf(w, "GAGAL CREATE ARTICLE TO REDIS ADA KESALAHAN DALAM PENYIMPANAN KE REDIS")
		return
	}

	// return response
	response, _ := json.Marshal(articleResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
