/**
 * @author mch
 */

package job

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-queue/rabbitmq"
	"queues/pkg/constant"
)

type ExecHandle func()
type Job struct {
	rabbitmqSend  rabbitmq.Sender
}

func NewJob(config Config) *Job {
	return &Job{
		rabbitmqSend: rabbitmq.MustNewSender(config.RabbitSenderConf),
	}
}

func(job *Job) Exec(ctx context.Context,handle ExecHandle,task *TaskDataMq) error {
	handle()
	byt,err := json.Marshal(task)
	if err != nil {
		return err
	}
	err = job.rabbitmqSend.Send("",constant.RabbitmqRouterKey,byt)
	return err
}

