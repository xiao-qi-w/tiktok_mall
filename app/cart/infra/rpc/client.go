package rpc

import (
	"sync"

	"github.com/cloudwego/biz-demo/gomall/common/clientsuite"

	"github.com/cloudwego/biz-demo/gomall/app/cart/conf"
	cartutils "github.com/cloudwego/biz-demo/gomall/app/cart/utils"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
	err           error
	registryAddr  string
	serviceName   string
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initProductClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartutils.MustHandleError(err)
}
