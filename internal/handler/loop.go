/**
 * @author mch
 */

package handler

import (
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/core/service"
	"queues/internal/svc"
	rabbitmqHandler "queues/internal/handler/rabbitmq"
)

type Loop struct {
	ServiceContext *svc.ServiceContext
}

func NewLoop(serviceContext *svc.ServiceContext) *Loop {
	return &Loop{ServiceContext: serviceContext}
}

func(l *Loop) Services() []service.Service {
	services := []service.Service{
		rabbitmq.MustNewListener(l.ServiceContext.Config.RabbitMqListenConf,rabbitmqHandler.NewRabbitMqHandler(l.ServiceContext)),
	}
	return services
}
