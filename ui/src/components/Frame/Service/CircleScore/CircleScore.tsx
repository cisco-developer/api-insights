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

import { HTMLAttributes } from 'react';
import { CircularProgress } from '@mui/material';
import ScoreLevel from '../ScoreLevel/ScoreLevel';
import './CircleScore.scss';

type Props = HTMLAttributes<HTMLElement> & {
  size?: number;
  thickness?: number;
  value?: number;
  darkTrack?: boolean;
  progress?: boolean;
};

export default function CircleScore(props: Props) {
  const size = props.size || 56;
  const thickness = props.thickness || 4;
  const { value } = props;

  const style = {
    width: `${size}px`,
    height: `${size}px`,
  };

  const scoreProgress = props.progress ? (
    <div className="score-progress">
      <CircularProgress
        color="inherit"
        variant="determinate"
        size={size}
        thickness={thickness}
        value={20}
      />
    </div>
  ) : null;

  return (
    <div className="circle-score" style={style}>
      <ScoreLevel score={value} className="circle-part">
        <div className={`score-track${props.darkTrack ? ' dark-track' : ''}`}>
          <CircularProgress
            color="inherit"
            variant="determinate"
            size={size}
            thickness={thickness}
            value={100}
          />
        </div>
        <div className="score-value">
          <CircularProgress
            color="inherit"
            variant="determinate"
            size={size}
            thickness={thickness}
            value={value || 0}
          />
        </div>
        {scoreProgress}
      </ScoreLevel>
      <div className="value-label">{props.children || value}</div>
    </div>
  );
}
