/**
 * @author mch
 */

package logic

import (
	"context"
	"fmt"
	"queues/internal/svc"
	"queues/pkg/job"
)

type KafkaService struct {
	ServiceContext *svc.ServiceContext
}
func NewKafkaService(serviceContext *svc.ServiceContext) *KafkaService {
	return &KafkaService{
		ServiceContext: serviceContext,
	}
}
func(s *KafkaService) Handle() func(ctx context.Context,data *job.TaskKafkaData) {
	return func(ctx context.Context, data *job.TaskKafkaData) {
		fmt.Println("kafka data:")
		fmt.Println(data)
	}
}
