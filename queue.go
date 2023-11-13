/**
 * @author mch
 */

package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"queues/internal/config"
	"queues/internal/handler"
	"queues/internal/svc"
)

var configFile = flag.String("f","etc/config.yaml","please enter file")
func main()  {
	flag.Parse()
	var c config.Config
	conf.MustLoad(*configFile,&c)
	if err := c.SetUp();err != nil {
		panic(err)
	}
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()
	serviceContext := svc.NewServiceContext(c)
	for _,s := range handler.NewLoop(serviceContext).Services() {
		serviceGroup.Add(s)
	}

	serviceGroup.Start()
}
