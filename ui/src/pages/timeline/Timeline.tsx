/*
 * Copyright 2022 Cisco Systems, Inc. and its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import { useEffect, useState } from 'react';
import { useOutletContext } from 'react-router-dom';
import Timeline from './components/Timeline/Timeline';
import { useFetchServiceCompliance } from '../../query/compliance';
import { AppFrameContext } from '../../components/Frame/AppFrame/AppFrame';
import {
  SpecData,
  useFetchSpecList,
} from '../../query/spec';

export default function TimelinePage() {
  const [timers] = useState({ refetchSpecTimerId: 0 });
  const { selectedService, refetchServiceList } = useOutletContext() as AppFrameContext;
  const {
    data: fullSpecList,
    isLoading: isSpecDataLoading,
    refetch: refetchSpecList,
  } = useFetchSpecList(selectedService?.id);
  const {
    data: complianceList,
    isLoading: isComplianceListLoading,
    refetch: refetchComplianceList,
  } = useFetchServiceCompliance(selectedService?.id);
  const specList = fullSpecList ? fullSpecList.filter(
    (spec: SpecData.Spec) => spec.state !== SpecData.SpecState.Reconstructed,
  ) : [];
  const snapshots = fullSpecList ? fullSpecList.filter(
    (spec: SpecData.Spec) => spec.state === SpecData.SpecState.Reconstructed,
  ) : [];
  const isSnapshotLoading = isSpecDataLoading;

  const refetchSpecCompliance = () => {
    refetchServiceList();
    refetchSpecList();
    refetchComplianceList();
  };

  const onReconstructed = () => {
    refetchSpecList();
  };

  useEffect(() => {
    if (!specList) return;

    // If a spec doesn't have a score, await backend analysis and perioidically refetch
    const needRefetchItem = specList.find(
      (i: SpecData.Spec) => i.score === null || i.score === undefined,
    );
    if (!needRefetchItem) return;

    if (timers.refetchSpecTimerId) {
      window.clearTimeout(timers.refetchSpecTimerId);
    }

    timers.refetchSpecTimerId = window.setTimeout(() => {
      refetchSpecCompliance();
    }, 2000);
  }, [specList]);

  const props = {
    selectedService,
    specList,
    isSpecDataLoading,
    complianceList,
    isComplianceListLoading,
    snapshots: snapshots as SpecData.Snapshot[],
    isSnapshotLoading,
    onSpecUploaded: refetchSpecCompliance,
    onReconstructed,
  };

  return <Timeline {...props} />;
}
