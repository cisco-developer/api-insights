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

import { useState, MouseEvent } from 'react';
import { Link } from 'react-router-dom';
import CircleScore from '../../../../components/Frame/Service/CircleScore/CircleScore';
import buildSortedListByStringField from '../../../../utils/order';
import TrendBarChart from '../TrendBarChart/TrendBarChart';
import SeveritySummary from '../../../../components/Severity/SeveritySummary/SeveritySummary';
import Revision, { RevisionData } from '../Revision/Revision';
import { ServiceData } from '../../../../query/service';
import './VersionReport.scss';

export type VersionData = {
  version: string;
  updated_at: string;
  revisions: RevisionData[];
  latestRevision: RevisionData;
};

type Props = {
  data: VersionData;
  service: ServiceData.Service;
  onReleased?: (item: RevisionData) => void;
  onArchived?: (item: RevisionData) => void;
};

export default function VersionReport(props: Props) {
  const [openDetail, setOpenDetail] = useState(false);
  const [showAllRevisions, setShowAllRevisions] = useState(false);
  const [highlightRevisionId, setHighLightRevisionId] = useState(null);

  const ascRevisionList = buildSortedListByStringField(
    props.data.revisions,
    'updated_at',
  );
  const lastRevision = props.data.revisions[0];

  const serviceNameId = props.service?.name_id || props.service?.id;

  const buildReportUrl = (specId: string) => `/reports?service=${serviceNameId}&spec=${specId}`;

  const onSetHighlightRevision = (item: RevisionData) => {
    setHighLightRevisionId(item.id);
  };

  const onCancelHighlightRevision = () => {
    setHighLightRevisionId(null);
  };

  const onEnterTrendBar = (e: MouseEvent<HTMLElement>) => {
    const { id } = e.currentTarget.dataset;
    setHighLightRevisionId(id);
  };

  const onViewMore = () => {
    setShowAllRevisions(!showAllRevisions);
  };

  const renderRevision = (data: RevisionData) => (
    <Revision
      key={data.id}
      className={data.id === highlightRevisionId ? 'highlight' : ''}
      data={data}
      linkTo={buildReportUrl(data.id)}
      onMouseEnter={onSetHighlightRevision}
      onMouseLeave={onCancelHighlightRevision}
      onReleased={props.onReleased}
      onArchived={props.onArchived}
    />
  );

  const renderRevisionList = () => {
    const DEFAULT_SHOW_COUNT = 4;
    const moreCount = props.data.revisions.length - DEFAULT_SHOW_COUNT;
    const showCount = Math.min(DEFAULT_SHOW_COUNT, props.data.revisions.length);
    const list = showAllRevisions
      ? props.data.revisions
      : props.data.revisions.slice(0, showCount);
    const viewMoreText = showAllRevisions
      ? 'View Less'
      : `View ${moreCount} More`;

    const viewMore = moreCount > 0 && (
      <div className="view-more-block">
        <div className="view-more" onClick={onViewMore}>
          {viewMoreText}
        </div>
      </div>
    );

    return (
      <div className="revision-list">
        {list.map(renderRevision)}
        {viewMore}
      </div>
    );
  };

  const renderTrends = () => {
    const data = ascRevisionList.map((i: RevisionData) => ({
      'data-id': i.id,
      label: i.revision,
      score: i.score,
      to: buildReportUrl(i.id),
      highlight: i.id === highlightRevisionId,
    }));

    return (
      <TrendBarChart
        data={data}
        minShareCount={7}
        onEnterBar={onEnterTrendBar}
        onLeaveBar={onCancelHighlightRevision}
      />
    );
  };

  const revisonList = renderRevisionList();
  const trends = renderTrends();
  const lastRevisionNoScore = lastRevision.score === undefined || lastRevision.score === null;

  const severitySummaryData = lastRevision.complianceList
    .filter((x) => x.analyzer !== 'drift')
    .map((i) => i.result.summary.stats);

  return (
    <div className="panel-block spec-version-report">
      <div className="item-brief" onClick={() => setOpenDetail(!openDetail)}>
        <div className="info-col">
          <div className="main-info">
            <div className="score-part">
              <CircleScore value={lastRevision.score} progress={lastRevisionNoScore} />
            </div>
            <div className="info-part">
              <div className="spec-info">
                <span className="service">{props.service.title}</span>
                <span className="version-label">{lastRevision.version}</span>
                <span className="revision-label">{lastRevision.revision}</span>
              </div>
            </div>
          </div>
          <div className="status-part block-item-light">
            <SeveritySummary
              data={severitySummaryData}
            />
            <Link to={buildReportUrl(lastRevision.id)} className="button-rc">
              View Full Report
            </Link>
          </div>
        </div>
        <div className="action-col" />
      </div>
      <div className={`item-detail${openDetail ? ' open' : ''}`}>
        <div className="detail-row">
          {revisonList}
          <div className="score-trends">
            <div className="trend-title-bar">
              <div className="title">Score Trends</div>
              <div className="title-tag">score per build</div>
            </div>
            {trends}
          </div>
        </div>
      </div>
    </div>
  );
}
