/*
Copyright 2018 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"managed-certs-gke/pkg/ingress"
)

type IngressController struct {
	client *ingress.Interface
	queue  workqueue.RateLimitingInterface
}

func (c *IngressController) Run(stopChannel <-chan struct{}, ingressWatcherDelay time.Duration) {
	defer c.queue.ShutDown()

	go c.runWatcher(ingressWatcherDelay)
	go wait.Until(c.synchronizeAllIngresses, time.Minute, stopChannel)

	<-stopChannel
}

func (c *IngressController) enqueue(obj interface{}) {
	if key, err := cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
	} else {
		c.queue.AddRateLimited(key)
	}
}

func (c *IngressController) synchronizeAllIngresses() {
	ingresses, err := c.client.List()
	if err != nil {
		runtime.HandleError(err)
		return
	}

	for _, ing := range ingresses.Items {
		c.enqueue(&ing)
	}
}
