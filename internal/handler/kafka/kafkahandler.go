/**
 * @author mch
 */

package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"queues/internal/logic"
	"queues/internal/svc"
	"queues/pkg/job"
)
type KfkHandle func(ctx context.Context,data *job.TaskKafkaData)
type KafkaHandler struct {
	ServiceContext *svc.ServiceContext
	KafkaMp map[string]KfkHandle
}
func NewKafkaHandler(serviceContext *svc.ServiceContext) *KafkaHandler {
	k := &KafkaHandler{ServiceContext: serviceContext}
	k.KafkaMp = map[string]KfkHandle{
		"test1": logic.NewKafkaService(serviceContext).Handle(),
	}
	return k
}

func(h *KafkaHandler) Consume(key, value string) error {
	fmt.Println("key=",key)
	fmt.Println("value:",value)
	data := []byte(value)
	var res job.TaskKafkaData
	err := json.Unmarshal(data,&res)
	if err != nil {
		return err
	}
	f := h.KafkaMp["test1"]
	f(context.Background(),&res)
	return nil
}
