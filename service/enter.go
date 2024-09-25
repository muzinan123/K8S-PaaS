package service

import (
	"kubeimooc.com/service/configmap"
	"kubeimooc.com/service/node"
	"kubeimooc.com/service/pod"
	"kubeimooc.com/service/secret"
)

//@Author: morris
type ServiceGroup struct {
	PodServiceGroup       pod.PodServiceGroup
	NodeServiceGroup      node.Group
	ConfigMapServiceGroup configmap.ServiceGroup
	SecretServiceGroup    secret.SeviceGroup
}

var ServiceGroupApp = new(ServiceGroup)
