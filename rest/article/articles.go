package article

import (
	"github.com/cisordeng/beego/xenon"

	bArticle "nature/business/article"
)

type Articles struct {
	xenon.RestResource
}

func init () {
	xenon.RegisterResource(new(Articles))
}

func (this *Articles) Resource() string {
	return "article.articles"
}

func (this *Articles) Params() map[string][]string {
	return map[string][]string{
		"GET":  []string{},
	}
}

func (this *Articles) Get() {
	page := this.GetPage()
	articles, pageInfo := bArticle.GetPagedArticles(page, xenon.Map{}, "-created_at")
	bArticle.Fill(articles)
	data := bArticle.EncodeManyArticle(articles)
	this.ReturnJSON(xenon.Map{
		"articles": data,
		"page_info": pageInfo.ToMap(),
	})
}
