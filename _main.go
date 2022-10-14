package main

import (
	"context"
	"fmt"
	"net/http"
	"try/go-router/domain/entity"
	"try/go-router/domain/repository"
	"try/go-router/mysql_repository"
	"try/go-router/pkg/mysql_connection"
	"try/go-router/pkg/redis_connection"
	"try/go-router/redis_repository"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	connectionDatabase     = mysql_connection.InitMysqlDB()
	articleRepositoryMysql = mysql_repository.NewArticleRepositoryMysqlInteractor(connectionDatabase)
	connectionRedis        = redis_connection.InitRedisClient()
	articleRepositoryRedis = redis_repository.NewRepoArticleRedisInteractor(connectionRedis)
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
	r.HandleFunc("/", ParamHandlerWithoutInput)
}

func insertData(ctx context.Context) *entity.Article {
	CreateFirstArticle := entity.DTONewCreateArticle{
		TitleOriginal: "Apa itu Lorem Ipsum?",
		TextOriginal:  "Lorem Ipsum hanyalah teks tiruan dari industri percetakan dan penyusunan huruf. Lorem Ipsum telah menjadi teks dummy standar industri sejak tahun 1500-an, ketika seorang pencetak yang tidak dikenal mengambil sekumpulan tipe dan mengacaknya untuk membuat buku spesimen tipe. Ini telah bertahan tidak hanya lima abad, tetapi juga lompatan ke pengaturan huruf elektronik, pada dasarnya tetap tidak berubah. Itu dipopulerkan pada 1960-an dengan merilis lembar Letraset yang berisi bagian-bagian Lorem Ipsum, dan baru-baru ini dengan perangkat lunak penerbitan desktop seperti Aldus PageMaker termasuk versi Lorem Ipsum.",
		Banner:        "example.jpg",
		Author:        "Taupik Pirdian",
		Thumbs:        "thumbs-example.jpg",
		IsHighLight:   false,
		Translation: []entity.DTOTranslation{
			{
				CodeLanguage: "ENG",
				Title:        "What is Lorem Ipsum?",
				Text:         "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
			},
			{
				CodeLanguage: "GER",
				Title:        "Was ist Lorem Ipsum?",
				Text:         "Lorem Ipsum ist einfach Blindtext der Druck- und Satzindustrie. Lorem Ipsum ist seit den 1500er Jahren der Standard-Dummy-Text der Branche, als ein unbekannter Drucker eine Reihe von Typen nahm und daraus ein Musterbuch für Typen erstellte. Sie hat nicht nur fünf Jahrhunderte, sondern auch den Sprung in den elektronischen Satz überstanden und ist im Wesentlichen unverändert geblieben. Es wurde in den 1960er Jahren mit der Veröffentlichung von Letraset-Blättern mit Passagen von Lorem Ipsum und in jüngerer Zeit mit Desktop-Publishing-Software wie Aldus PageMaker, einschließlich Versionen von Lorem Ipsum, populär.",
			},
		},
	}

	// create dan generate id code number dan year modeling
	FirstArticle, errCheckDomainArticle := entity.NewCreateArticle(CreateFirstArticle)
	if errCheckDomainArticle != nil {
		fmt.Println("GAGAL CREATE ARTICLE KARENA WRONG DOMAIN")
		panic(errCheckDomainArticle)

	}

	fmt.Println("=> => => Process Store Data To DB")
	handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
	errStoreRepo := handlerRepo.articleRepository.Store(ctx, FirstArticle)
	if errStoreRepo != nil {
		fmt.Println("GAGAL CREATE ARTICLE ADA KESALAHAN DALAM PENYIMPANAN")
		panic(errStoreRepo)
	}
	fmt.Println("... Success Store Data To DB")

	// store to redis
	handlerRepoRedis := NewArticleLogicRedisFactoryHandler(articleRepositoryRedis)
	errStoreRepoRedis := handlerRepoRedis.articleRepository.StoreOrUpdateData(ctx, FirstArticle)
	if errStoreRepoRedis != nil {
		fmt.Println("GAGAL CREATE ARTICLE TO REDIS ADA KESALAHAN DALAM PENYIMPANAN KE REDIS")
		panic(errStoreRepoRedis)
	}
	fmt.Println("... Success Store Data To Redis")

	return FirstArticle
}

func selectDataByCode(ctx context.Context, codeArticle string) (*entity.Article, error) {
	handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
	dataArticle, errGetData := handlerRepo.articleRepository.GetDataByCode(ctx, codeArticle)
	if errGetData != nil {
		panic(errGetData)
	}

	if dataArticle != nil {
		fmt.Println("Kode : " + dataArticle.GetCodeArtikel() + ", Judul : " + dataArticle.GetTitleArtikel() + ", Author : " + dataArticle.GetAuthorArtikel())
	}

	return dataArticle, errGetData
}

func setDataToRedisByCode(ctx context.Context, result *entity.Article) {
	// store to redis
	handlerRepoRedis := NewArticleLogicRedisFactoryHandler(articleRepositoryRedis)
	errStoreRepoRedis := handlerRepoRedis.articleRepository.StoreOrUpdateData(ctx, result)
	if errStoreRepoRedis != nil {
		fmt.Println("GAGAL CREATE ARTICLE TO REDIS ADA KESALAHAN DALAM PENYIMPANAN KE REDIS")
		panic(errStoreRepoRedis)
	} else {
		fmt.Println("... Success Store Single Data TO Redis")
	}
}

func selectDataFromRedisByCode(ctx context.Context, codeArticle string) {
	fmt.Println("=> => => Process Select Data From Redis")
	handlerRepo := NewArticleLogicRedisFactoryHandler(articleRepositoryRedis)
	result, errGetData := handlerRepo.articleRepository.GetAttributeArticleByKode(ctx, codeArticle)
	if errGetData != nil {
		panic(errGetData)
	}

	if result == nil {
		fmt.Println("=> => => Process Get Single Data From DB with Code: " + codeArticle)
		resultGetDb, _ := selectDataByCode(ctx, codeArticle)
		if resultGetDb != nil {
			fmt.Println("... Set Single Data To Redis with Code: " + codeArticle)
			setDataToRedisByCode(ctx, resultGetDb)
		}
	} else {
		fmt.Println("... Success Get Single Data From Redis with Code: " + codeArticle)
		fmt.Println("Kode : " + result.GetCodeArtikel() + ", Judul : " + result.GetTitleArtikel() + ", Author : " + result.GetAuthorArtikel())
	}
}

func selectAllData(ctx context.Context) {
	handlerRepo := NewArticleLogicFactoryHandler(articleRepositoryMysql)
	dataCollectionArticle, errGetData := handlerRepo.articleRepository.GetAllData(ctx)
	if errGetData != nil {
		panic(errGetData)
	}

	for _, dataArticle := range dataCollectionArticle {
		fmt.Println("Kode : " + dataArticle.GetCodeArtikel() + ", Judul : " + dataArticle.GetTitleArtikel() + ", Author : " + dataArticle.GetAuthorArtikel())
	}
}

func selectAllDataFromRedis(ctx context.Context) {
	handlerRepo := NewArticleLogicRedisFactoryHandler(articleRepositoryRedis)
	result, errGetData := handlerRepo.articleRepository.GetAttributeArticleByKode(ctx, "XXX")
	if errGetData != nil {
		panic(errGetData)
	}

	if result == nil {
		fmt.Println("... Get Data From DB and Set To Redis")
		selectAllData(ctx)
	} else {
		fmt.Println("... Get Data From Redis")

		// for _, dataArticle := range dataCollectionArticle {
		// 	fmt.Println("Kode : " + dataArticle.GetCodeArtikel() + ", Judul : " + dataArticle.GetTitleArtikel() + ", Author : " + dataArticle.GetAuthorArtikel())
		// }
	}
}

func checkConnectionRedis() {
	fmt.Println("Go Redis Connection Test")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	result := ping(client)
	if result == nil {
		fmt.Println("Redis Connected")
	} else {
		fmt.Println("Redis Not Connected")
	}
}

// check koneksi
func ping(client *redis.Client) error {
	pong, err := client.Ping(client.Context()).Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	// Output: PONG <nil>

	return nil
}

/*
function route
*/
func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	fmt.Println("OK")
	// fmt.Fprintf("OK")
}
