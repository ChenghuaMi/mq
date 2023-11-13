/**
 * @author mch
 */

package rabbitmq

import (
	"context"
	"encoding/json"
	"queues/internal/logic"
	"queues/internal/svc"
	"queues/pkg/job"
)
type MqHandler func(ctx context.Context,data *job.TaskDataMq)

type RabbitMqHandler struct {
	ServiceContext *svc.ServiceContext
	DataMp map[string]MqHandler
}

func NewRabbitMqHandler(serviceContext *svc.ServiceContext) *RabbitMqHandler {
	handle := &RabbitMqHandler{
		ServiceContext: serviceContext,
	}
	handle.DataMp = map[string]MqHandler{
		"test1": logic.NewLogicService(serviceContext).Do(),
	}
	return handle
}
func(r *RabbitMqHandler) Consume(message string) error {
	var task job.TaskDataMq
	data := []byte(message)
	json.Unmarshal(data,&task)
	h,ok := r.DataMp["test1"]
	if ok {
		h(context.Background(),&task)
	}
	return nil
}