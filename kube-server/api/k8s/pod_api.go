package k8s

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"kubeimooc.com/global"
	pod_req "kubeimooc.com/model/pod/request"
	"kubeimooc.com/response"
)

/*
k8s.io/api/apps/v1 ：对应K8S apiVersion: apps/v1接口操作的对象
如Deployment、DaemonSet、StatefulSet等
k8s.io/api/core/v1 ：对应K8S apiVersion: v1接口操作的对象
如：ConfigMap、Service、NameSpace等
k8s.io/apimachinery/pkg/apis/meta/v1：对对象的实际操作，如增删改查等。
k8s.io/client-go/kubernetes 用于链接k8s集群
*/

type PodApi struct {
}

// 因为update的字段属性有限 而我们实际更新过程当中 会修改定义的任一字段
func UpdatePod(ctx context.Context, pod *corev1.Pod) error {
	//update
	_, err := global.KubeConfigSet.CoreV1().Pods(pod.Namespace).Update(ctx, pod, metav1.UpdateOptions{})
	return err
}

func PatchPod(patchData map[string]interface{}, k8sPod *corev1.Pod, ctx context.Context) error {
	patchDataBytes, _ := json.Marshal(&patchData)
	_, err := global.KubeConfigSet.CoreV1().Pods(k8sPod.Namespace).Patch(
		ctx,
		k8sPod.Name,
		types.StrategicMergePatchType,
		patchDataBytes,
		metav1.PatchOptions{},
	)
	return err
}

func (*PodApi) DeletePod(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")
	err := podService.DeletePod(namespace, name)
	if err != nil {
		response.FailWithMessage(c, "删除Pod失败，detail："+err.Error())
	} else {
		response.Success(c)
	}
}

func (*PodApi) CreateOrUpdatePod(c *gin.Context) {
	var podReq pod_req.Pod
	if err := c.ShouldBind(&podReq); err != nil {
		response.FailWithMessage(c, "参数解析失败，detail："+err.Error())
		return
	}
	//校验必填项
	if err := podValidate.Validate(&podReq); err != nil {
		response.FailWithMessage(c, "参数验证失败，detail："+err.Error())
		return
	}
	if msg, err := podService.CreateOrUpdatePod(podReq); err != nil {
		response.FailWithMessage(c, msg)
	} else {
		response.SuccessWithMessage(c, msg)
	}
}

func (*PodApi) GetPodListOrDetail(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Query("name")
	keyword := c.Query("keyword")
	if name != "" {
		detail, err := podService.GetPodDetail(namespace, name)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "获取Pod详情成功", detail)
	} else {
		err, items := podService.GetPodList(namespace, keyword, c.Query("nodeName"))
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		response.SuccessWithDetailed(c, "获取Pod列表成功", items)
	}
}
