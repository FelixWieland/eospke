package config

import (
	"fmt"
	"time"

	elastic_logrus "github.com/interactive-solutions/go-logrus-elasticsearch"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

const elasticsearchHost = "http://localhost:9200"
const elasticsearchUsername = "elastic"
const elasticsearchPassword = "changeme"

func MakeElasticHook(logger *logrus.Logger) (*elastic_logrus.ElasticSearchHook, error) {

	elastic.SetSniff(false)

	url := elastic.SetURL(elasticsearchHost)
	auth := elastic.SetBasicAuth(elasticsearchUsername, elasticsearchPassword)

	client, err := elastic.NewClient(url, auth)

	if err != nil {
		logger.WithError(err).Fatal("Failed to construct elasticsearch client")
	}
	hook, err := elastic_logrus.NewElasticHook(client, elasticsearchHost, logrus.DebugLevel, func() string {
		return fmt.Sprintf("%s-%s", "some-index", time.Now().Format("2006-01-02"))
	}, time.Second*15)

	return hook, err
}
