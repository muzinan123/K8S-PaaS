package configmap

import "kubeimooc.com/convert"

//@Author: morris
type ServiceGroup struct {
	ConfigMapService
}

var configConvert = convert.ConvertGroupApp.ConfigMapConvert
