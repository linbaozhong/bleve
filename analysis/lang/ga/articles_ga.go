package ga

import (
	"github.com/linbaozhong/bleve/v2/analysis"
	"github.com/linbaozhong/bleve/v2/registry"
)

const ArticlesName = "articles_ga"

// this content was obtained from:
// lucene-4.7.2/analysis/common/src/resources/org/apache/lucene/analysis

var IrishArticles = []byte(`
d
m
b
`)

func ArticlesTokenMapConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.TokenMap, error) {
	rv := analysis.NewTokenMap()
	err := rv.LoadBytes(IrishArticles)
	return rv, err
}

func init() {
	registry.RegisterTokenMap(ArticlesName, ArticlesTokenMapConstructor)
}
