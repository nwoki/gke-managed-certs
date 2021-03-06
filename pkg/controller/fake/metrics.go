/*
Copyright 2020 Google LLC

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

package fake

import (
	"time"

	"github.com/GoogleCloudPlatform/gke-managed-certs/pkg/controller/metrics"
)

type FakeMetrics struct {
	ManagedCertificatesStatuses           map[string]int
	SslCertificateBackendErrorObserved    int
	SslCertificateQuotaErrorObserved      int
	SslCertificateBindingLatencyObserved  int
	SslCertificateCreationLatencyObserved int
}

var _ metrics.Metrics = &FakeMetrics{}

func NewMetrics() *FakeMetrics {
	return &FakeMetrics{}
}

func NewMetricsSslCertificateCreationAlreadyReported() *FakeMetrics {
	metrics := NewMetrics()
	metrics.SslCertificateCreationLatencyObserved++
	return metrics
}

func (f *FakeMetrics) Start(address string) {}

func (f *FakeMetrics) ObserveManagedCertificatesStatuses(statuses map[string]int) {
	f.ManagedCertificatesStatuses = statuses
}

func (f *FakeMetrics) ObserveSslCertificateBackendError() {
	f.SslCertificateBackendErrorObserved++
}

func (f *FakeMetrics) ObserveSslCertificateQuotaError() {
	f.SslCertificateQuotaErrorObserved++
}

func (f *FakeMetrics) ObserveSslCertificateBindingLatency(creationTime time.Time) {
	f.SslCertificateBindingLatencyObserved++
}

func (f *FakeMetrics) ObserveSslCertificateCreationLatency(creationTime time.Time) {
	f.SslCertificateCreationLatencyObserved++
}
