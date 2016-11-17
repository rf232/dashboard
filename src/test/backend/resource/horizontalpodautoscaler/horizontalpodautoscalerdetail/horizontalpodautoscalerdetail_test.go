// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package horizontalpodautoscalerdetail

import (
	"reflect"
	"testing"

	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/kubernetes/dashboard/src/app/backend/resource/horizontalpodautoscaler"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/autoscaling"
	"k8s.io/kubernetes/pkg/client/unversioned/testclient"
)

// func GetHorizontalPodAutoscalerDetail(client *client.Client, namespace string, name string) (*HorizontalPodAutoscalerDetail, error) 

func TestGetHorizontalPodAutoscalerDetail(t *testing.T) {
	cases := []struct {
		namespace, name string
		expectedActions []string
		hpa             *autoscaling.HorizontalPodAutoscaler
		expected        *HorizontalPodAutoscalerDetail
	}{
		{
			"test-namespace", "test-name",
			[]string{"get"},
			&autoscaling.HorizontalPodAutoscaler{
				ObjectMeta: api.ObjectMeta{Name: "test-name"},
				Spec: autoscaling.HorizontalPodAutoscalerSpec{
					ScaleTargetRef: autoscaling.CrossVersionObjectReference{
						Kind: "test-kind",
						Name: "test-name2",
					},
					MaxReplicas: 3,
				},
				Status: autoscaling.HorizontalPodAutoscalerStatus{
					CurrentReplicas: 1,
					DesiredReplicas: 2,
				},
			},
			&HorizontalPodAutoscalerDetail{
				ObjectMeta:     common.ObjectMeta{Name: "test-name"},
				TypeMeta:       common.TypeMeta{Kind: common.ResourceKindHorizontalPodAutoscaler},
				ScaleTargetRef: horizontalpodautoscaler.ScaleTargetRef{
					Kind: "test-kind",
					Name: "test-name2",
				},
				MaxReplicas:     3,
				CurrentReplicas: 1,
				DesiredReplicas: 2,
			},
		},
	}

	for _, c := range cases {
		fakeClient := testclient.NewSimpleFake(c.hpa)

		actual, _ := GetHorizontalPodAutoscalerDetail(fakeClient, c.namespace, c.name)

		actions := fakeClient.Actions()
		if len(actions) != len(c.expectedActions) {
			t.Errorf("Unexpected actions: %v, expected %d actions got %d", actions,
				len(c.expectedActions), len(actions))
			continue
		}

		for i, verb := range c.expectedActions {
			if actions[i].GetVerb() != verb {
				t.Errorf("Unexpected action: %+v, expected %s",
					actions[i], verb)
			}
		}

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("GetEvents(client,heapsterClient,%#v, %#v) == \ngot: %#v, \nexpected %#v",
				c.namespace, c.name, actual, c.expected)
		}
	}
}