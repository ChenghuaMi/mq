/**
 * @author mch
 */

package job

import (
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
)

type KfkJob struct {
	KfkSend *kq.Pusher
}

func NewKfkJob(config KfkConfig) *KfkJob {
	return &KfkJob{
		KfkSend: kq.NewPusher(config.Addr,config.Topic),
	}
}
func(l *KfkJob) Exec(data TaskKafkaData) {
	res,_ := json.Marshal(data)
	l.KfkSend.Push(string(res))
}