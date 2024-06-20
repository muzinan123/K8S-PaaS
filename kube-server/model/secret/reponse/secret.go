package request

import (
	corev1 "k8s.io/api/core/v1"
	"kubeimooc.com/model/base"
)

//@Author: morris
type Secret struct {
	Name      string             `json:"name"`
	Namespace string             `json:"namespace"`
	DataNum   int                `json:"dataNum"`
	Age       int64              `json:"age"`
	Type      corev1.SecretType  `json:"type"`
	Labels    []base.ListMapItem `json:"labels"`
	Data      []base.ListMapItem `json:"data"`
}
