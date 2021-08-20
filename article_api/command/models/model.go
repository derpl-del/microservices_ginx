package models

import (
	"article_api/command/config/pgadapter"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"
)

type Article struct {
	ID      uint       `json:"id" gorm:"not null;primaryKey"`
	Author  string     `json:"author"`
	Title   string     `json:"title" gorm:"type:text"`
	Body    string     `json:"body" gorm:"type:text"`
	Created *time.Time `json:"created"`
}

type Response struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func OpenConection() pgadapter.Adapter {
	Adapter := pgadapter.Adapter{}.New()
	return Adapter
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

func (A *Article) Send() error {
	url := os.Getenv("API_ENV")
	url = url + "/api/v1/create"
	jsonStr, _ := json.Marshal(A)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
