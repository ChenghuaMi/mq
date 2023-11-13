/**
 * @author mch
 */

package svc

import "queues/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(config config.Config) *ServiceContext {
	return &ServiceContext{Config: config}
}