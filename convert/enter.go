package convert

import (
	"kubeimooc.com/convert/configmap"
	"kubeimooc.com/convert/node"
	"kubeimooc.com/convert/pod"
	"kubeimooc.com/convert/secret"
)

//@Author: morris

type ConvertGroup struct {
	PodConvert       pod.PodConvertGroup
	NodeConvert      node.Group
	ConfigMapConvert configmap.ConvertGroup
	SecretConvert    secret.ConvertGroup
}

var ConvertGroupApp = new(ConvertGroup)
