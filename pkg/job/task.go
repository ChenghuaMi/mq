/**
 * @author mch
 */

package job

// TaskDataMq 投递的数据格式
type TaskDataMq struct {
	UserName string
	Data interface{}
}

type TaskKafkaData struct {
	Id int
	Username string
	Age int
}