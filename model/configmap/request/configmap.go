package request

import "kubeimooc.com/model/base"

//@Author: morris

type ConfigMap struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	Labels    []base.ListMapItem `json:"labels"`
	Data      []base.ListMapItem `json:"data"`
}
