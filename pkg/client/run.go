package client

import (
	"fmt"
	"time"

	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
)

// Watcher
type Watcher struct {
	kube       *kubernetes.Clientset
	queue      workqueue.RateLimitingInterface
	store      cache.Store
	controller cache.Controller
}

// Init a watcher instance
func (w Watcher) Init() {
	var err error
	if w.kube, err = initializeKubeClient(""); err != nil {
		fmt.Println("initializeKubeClient err:", err)
	}

	w.queue = workqueue.NewRateLimitingQueue(workqueue.NewItemExponentialFailureRateLimiter(
		100*time.Millisecond,
		5*time.Second,
	))

	w.store, w.controller = cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return w.kube.CoreV1().Secrets(core_v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return w.kube.CoreV1().Secrets(core_v1.NamespaceAll).Watch(options)
			},
		},
		&core_v1.Secret{},
		0,
		cache.ResourceEventHandlerFuncs{
			AddFunc:    w.hostsEventHandlerAdd,
			UpdateFunc: w.hostsEventHandlerUpdate,
			DeleteFunc: w.hostsEventHandlerDelete,
		},
	)
}

// build Kubernetes Clientset from kubeconfig, or fallback to in-cluster initialization
// if kubeconfigPath is empty
func initializeKubeClient(kubeconfigPath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

// Run runs the event loop
// stopCh channel is used to send interrupt signal to stop it.
func (w Watcher) Run(stopCh <-chan struct{}) {
	fmt.Println("client/run:Running")
	go w.controller.Run(stopCh)
}

// hostsEventHandlerAdd
func (w Watcher) hostsEventHandlerAdd(object interface{}) {
	fmt.Println("hostsEventHandlerAdd")
	fmt.Println(object)
}

// hostsEventHandlerUpdate
func (w Watcher) hostsEventHandlerUpdate(oldObject, newObject interface{}) {
	fmt.Println("hostsEventHandlerUpdate")
	fmt.Println(oldObject, newObject)
}

// hostsEventHandlerDelete
func (w Watcher) hostsEventHandlerDelete(object interface{}) {
	fmt.Println("hostsEventHandlerDelete")
	fmt.Println(object)
}
