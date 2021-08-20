package pipeline

import (
	"article_api/query/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

var c = cache.New(1*time.Minute, 3*time.Minute)

func GetArticle(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var Search models.Search
	json.Unmarshal(ReqBody, &Search)
	res, found := Search.GetCache(c)
	if !found {
		var Articles []models.Article
		err := Search.GetQuery(&Articles)
		if err != nil {
			res.ErrorCode = "9999"
			res.ErrorMessage = fmt.Sprint(err)
		} else {
			res.ErrorCode = "0000"
			res.ErrorMessage = "success"
			res.Articles = Articles
			Search.CreateCache(c, &res)
		}
	}
	json.NewEncoder(w).Encode(res)
}

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var Article models.Article
	json.Unmarshal(ReqBody, &Article)
	Article.Create()
	w.WriteHeader(200)
}
