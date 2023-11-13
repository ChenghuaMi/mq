/**
 * @author mch
 */

package logic

import (
	"fmt"
	"context"
	//"queues/internal/handler/rabbitmq"
	"queues/internal/svc"
	"queues/pkg/job"
)

type LogicService struct {
	ServiceCtx *svc.ServiceContext
}
func NewLogicService(serviceCtx *svc.ServiceContext) *LogicService {
	return &LogicService{ServiceCtx: serviceCtx}
}
func(l *LogicService) Do() func(ctx context.Context,data *job.TaskDataMq) {
	return func(ctx context.Context, data *job.TaskDataMq) {
		fmt.Println("doiiiiiiii")
		fmt.Println("data:",data)
	}
}
