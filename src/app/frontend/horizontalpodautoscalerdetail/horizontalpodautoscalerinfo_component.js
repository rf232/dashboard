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

import {StateParams} from 'common/resource/resourcedetail';
import {stateName as deploymentStateName} from 'deploymentdetail/deploymentdetail_state';
import {stateName as replicaSetStateName} from 'replicasetdetail/replicasetdetail_state';
import {stateName as replicationControllerStateName} from 'replicationcontrollerdetail/replicationcontrollerdetail_state';

const referenceKindToDetailStateName = {
  Deployment: deploymentStateName,
  ReplicaSet: replicaSetStateName,
  ReplicationController: replicationControllerStateName,
};

/**
 * @final
 */
export default class HorizontalPodAutoscalerInfoController {
  /**
   * Constructs horizontal pod autoscaler controller info object.
   *
   * @param {!ui.router.$state} $state
   * @param {!angular.$interpolate} $interpolate
   * @ngInject;
   */
  constructor($state, $interpolate) {
    /** @private {!ui.router.$state} */
    this.state_ = $state;

    /** @private {!angular.$interpolate} */
    this.interpolate_ = $interpolate;

    /**
     * Horizontal Pod Autoscaler details. Initialized from the scope.
     * @export {!backendApi.HorizontalPodAutoscalerDetail}
     */
    this.horizontalPodAutoscaler;
  }

  /**
   * @return {string}
   * @export
   */
  getScaleTargetHref() {
    return this.state_.href(
        referenceKindToDetailStateName[this.horizontalPodAutoscaler.scaleTargetRef.kind],
        new StateParams(
            this.horizontalPodAutoscaler.objectMeta.namespace,
            this.horizontalPodAutoscaler.scaleTargetRef.name));
  }
  /**
   * @export
   * @return {string} localized tooltip with the formatted creation date
   */
  getLatScaledTooltip() {
    let filter = this.interpolate_(`{{date | date}}`);
    /** @type {string} @desc Tooltip 'Last scaled at [some date]' showing the exact time of
     * the last time the horizontal pod autoscaler was scaled.*/
    let MSG_HORIZONTAL_POD_AUTOSCALER_DETAIL_LAST_SCALED_TOOLTIP = goog.getMsg(
        'Last scaled at {$scaleDate}',
        {'scaleDate': filter({'date': this.horizontalPodAutoscaler.lastScaleTime})});
    return MSG_HORIZONTAL_POD_AUTOSCALER_DETAIL_LAST_SCALED_TOOLTIP;
  }
}

/**
 * Definition object for the component that displays horizontal pod autoscaler info.
 *
 * @return {!angular.Directive}
 */
export const horizontalPodAutoscalerInfoComponent = {
  controller: HorizontalPodAutoscalerInfoController,
  templateUrl: 'horizontalpodautoscalerdetail/horizontalpodautoscalerinfo.html',
  bindings: {
    /** {!backendApi.HorizontalPodAutoscalerDetail} */
    'horizontalPodAutoscaler': '=',
  },
};
