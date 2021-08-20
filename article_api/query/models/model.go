package models

import (
	"article_api/query/config/pgadapter"
	"errors"
	"time"

	"github.com/patrickmn/go-cache"
)

type Search struct {
	Query  string `json:"query"`
	Author string `json:"author"`
}

type Article struct {
	ID      uint       `json:"id" gorm:"not null;primaryKey"`
	Author  string     `json:"author"`
	Title   string     `json:"title" gorm:"type:text"`
	Body    string     `json:"body" gorm:"type:text"`
	Created *time.Time `json:"created"`
}

type Response struct {
	ErrorCode    string    `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Articles     []Article `json:"list_article"`
}

func OpenConection() pgadapter.Adapter {
	Adapter := pgadapter.Adapter{}.New()
	return Adapter
}

func (S *Search) GetQuery(list *[]Article) error {
	adp := OpenConection()
	query := "%" + S.Query + "%"
	err := adp.Table.Where("body LIKE ? or title like ?", query, query).Find(&list, &Article{Author: S.Author}).Error
	adp.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (A *Article) Create() error {
	err := A.IsValid()
	if err != nil {
		return err
	}
	adp := OpenConection()
	times := time.Now()
	A.Created = &times
	err = adp.Table.Create(A).Error
	adp.Connection.Close()
	if err != nil {
		return err
	}
	return nil
}

func (A *Article) IsValid() (err error) {
	if len(A.Author) < 1 {
		return errors.New("article author cannot be null")
	} else if len(A.Title) < 1 {
		return errors.New("article title cannot be null")
	} else if len(A.Body) < 1 {
		return errors.New("article body cannot be null")
	}
	return
}

func (S *Search) GetCache(c *cache.Cache) (Response, bool) {
	tittle := S.Author + "|" + S.Query
	foo, found := c.Get(tittle)
	if found {
		art := foo.(*Response)
		return *art, found
	}
	return Response{}, found
}

func (S *Search) CreateCache(c *cache.Cache, res *Response) {
	tittle := S.Author + "|" + S.Query
	c.Set(tittle, res, cache.DefaultExpiration)
}
