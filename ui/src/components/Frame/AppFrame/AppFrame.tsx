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
import { Outlet, useSearchParams } from 'react-router-dom';
import { ServiceData, useFetchServiceList } from '../../../query/service';
import { InfoData, useFetchInfo } from '../../../query/info';
import AppHeader from '../AppHeader/AppHeader';
import './AppFrame.scss';

export type AppFrameContext = {
  sysInfo: InfoData.Info;
  navStore: { [index: string]: string };
  serviceList: ServiceData.Service[];
  selectedService: ServiceData.Service;
  specServiceSummary: ServiceData.ServiceSummary;
  onServiceSelected: (e: ServiceData.Service) => void;
  refetchServiceList: () => void;
  updateSpecSummary: (e: ServiceData.ServiceSummary) => void;
};

type Props = {
  navStore: {
    [index: string]: string;
  };
};

export default function AppFrame(props: Props) {
  const { data: sysInfo } = useFetchInfo();
  const { data: serviceList, refetch: refetchServiceList } = useFetchServiceList();
  const [searchParams, setSearchParams] = useSearchParams();
  const [specServiceSummary, setSpecServiceSummary] = useState(null);
  const serviceId = searchParams.get('service');
  const [navStore, setNavStore] = useState(props.navStore);
  const selectedService = (serviceList || []).find(
    (element: ServiceData.Service) => element.name_id === serviceId || element.id === serviceId,
  );
  const serviceSearchParam = `service=${serviceId}`;
  const otherSearchParam = searchParams
    .toString()
    .replace(serviceSearchParam, '');
  const newNavStore = { ...navStore };

  const onServiceSelected = (e: ServiceData.Service) => {
    setSearchParams({ service: e.name_id || e.id });
    setSpecServiceSummary(null);
  };

  const updateSpecSummary = (e: ServiceData.ServiceSummary) => {
    setSpecServiceSummary(e);
  };

  useEffect(() => {
    Object.keys(navStore).forEach((key) => {
      newNavStore[key] = serviceSearchParam;
    });
    setNavStore(newNavStore);
    setSpecServiceSummary(null);
  }, [serviceId]);

  useEffect(() => {
    const page = window.location.pathname.substring(
      window.location.pathname.lastIndexOf('/') + 1,
    );

    if (newNavStore[page] !== undefined) {
      newNavStore[page] = searchParams.toString();
      setNavStore(newNavStore);
    }
  }, [otherSearchParam]);

  const context: AppFrameContext = {
    sysInfo: sysInfo || {},
    navStore,
    serviceList,
    selectedService,
    specServiceSummary,
    refetchServiceList,
    onServiceSelected,
    updateSpecSummary,
  };

  return (
    <div className="app">
      <AppHeader />
      <div className="app-body">
        <Outlet context={context} />
      </div>
    </div>
  );
}
