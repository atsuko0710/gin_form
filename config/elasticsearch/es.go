package elasticsearch

import (
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
)

var Client *elastic.Client

func Init() (err error) {
	Client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(viper.GetString("es.url")))
	if err != nil {
		return err
	}

	return nil
}
