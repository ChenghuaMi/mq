/**
 * @author mch
 */

package job

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-queue/rabbitmq"
	"testing"
)

func TestRabbitmqSend(t *testing.T) {
	cf := rabbitmq.RabbitSenderConf{
	}
	cf.Username = "admin"
	cf.Password = "admin"
	cf.Port = 5672
	cf.Host = "192.168.2.115"
	cf.VHost = "/"
	cfg := Config{
		RabbitSenderConf:cf,
	}
	NewJob(cfg).Exec(context.Background(), func() {
		fmt.Println("start send.....")
	},&TaskDataMq{
		UserName: "mch",
		Data:     "this is a rabbitmq ",
	})
}
