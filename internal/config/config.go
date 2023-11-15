/**
 * @author mch
 */

package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-queue/rabbitmq"
	"github.com/zeromicro/go-zero/rest"
)
type Config struct {
	rest.RestConf
	RabbitMqListenConf rabbitmq.RabbitListenerConf
	KafkaListenConf kq.KqConf
}

