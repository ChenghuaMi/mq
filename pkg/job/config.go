/**
 * @author mch
 */

package job

import "github.com/zeromicro/go-queue/rabbitmq"

type Config struct {
	RabbitSenderConf rabbitmq.RabbitSenderConf
}
