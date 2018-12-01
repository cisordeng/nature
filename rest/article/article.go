package article

import (
	b_article "mango/business/article"
	"mango/rest"
)

type Article struct {
	rest.RestResource
}

func (o *Article) Get() {
	articleId, _ := o.GetInt("article_id", 0)
	article, err, be := b_article.GetArticleById(articleId)
	data := b_article.EncodeArticle(article)
	response := rest.MakeResponse(data, err, be)
	o.ReturnJSON(response)
}

func (o *Article) GetAll() {
}


func (o *Article) Put() {
	Title := o.GetString("title", "")
	Content := o.GetString("content", "")
	article, err, be := b_article.Create(Title, Content)
	data := b_article.EncodeArticle(article)
	response := rest.MakeResponse(data, err, be)
	o.ReturnJSON(response)
}


