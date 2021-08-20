package pipeline

import (
	"article_api/command/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var res models.Response

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var Article models.Article
	json.Unmarshal(ReqBody, &Article)
	err := Article.Create()
	if err != nil {
		res.ErrorCode = "9999"
		res.ErrorMessage = fmt.Sprint(err)
	} else {
		res.ErrorCode = "0000"
		res.ErrorMessage = "success"
		Article.Send()
	}
	json.NewEncoder(w).Encode(res)
}
