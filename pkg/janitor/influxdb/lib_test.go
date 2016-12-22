package influx

import (
	"fmt"
	"net/url"
	"testing"
	api "github.com/appscode/api/kubernetes/v1beta1"
	influxdb "github.com/influxdata/influxdb/client"
	"github.com/stretchr/testify/assert"
)

func TestInfluxJanitor(t *testing.T) {
	host := ""
	port := ""
	user := ""
	pass := ""
	u, _ := url.Parse(fmt.Sprintf("http://%v:%v", host, port))
	iConfig := influxdb.Config{
		URL:       *u,
		Username:  user,
		Password:  pass,
		UserAgent: fmt.Sprintf("%v/%v", "kubed", 1.0),
	}

	// InfluxDB client
	influxClient, err := influxdb.NewClient(iConfig)
	assert.Nil(t, err)

	settings := &api.ClusterSettings{
		MonitoringStorageLifetime: 6 * 60 * 60,
	}
	err = UpdateRetentionPolicy(influxClient, settings)
	assert.Nil(t, err)
}
