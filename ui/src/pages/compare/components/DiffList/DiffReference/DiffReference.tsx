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

import { PropaneSharp } from '@mui/icons-material';
import { editor } from 'monaco-editor';
import { useState, useEffect } from 'react';
import { MonacoDiffEditor } from 'react-monaco-editor';
import { DiffData } from '../../../../../query/compare';
import { SpecData } from '../../../../../query/spec';
import iterateObject from '../../../../../utils/iterateObject';
import { monacoMount } from '../../../../../utils/monacoInjection';
import waitFor from '../../../../../utils/waitFor';
import './DiffReference.scss';

type Props = {
    data: string;
    leftSpec?: SpecData.Spec;
    rightSpec?: SpecData.Spec;
    removeActive: (ref: string) => void;
    setActiveRefs: (refs: string[]) => void;
};

export default function DiffReference(props: Props) {
  const leftSchemaSection = (props.leftSpec)
    ? JSON.parse(props.leftSpec?.doc) : {};
  const rightSchemaSection = (props.rightSpec)
    ? JSON.parse(props.rightSpec?.doc) : {};
  const refList = props.data.split('/');
  const findSchema = (refNavList: string[], doc = '') => {
    let current = JSON.parse(doc);
    try {
      refNavList.forEach((ref) => {
        if (ref !== '#') {
          current = current[ref];
        }
      });
    } catch (_) {
      return undefined;
    }
    return current;
  };
  const refName = props.data.split('/schemas/')[1];
  const [refs, setRefs] = useState<string[]>([]);
  useEffect(() => {
    iterateObject(leftSchemaSection, refs);
    iterateObject(rightSchemaSection, refs);
  }, []);
  return (
    <div id={refName}>
      <div className="reference-button" onClick={() => props.removeActive(props.data)}>
        <div className="reference-label">Viewing reference</div>
        <div className="reference-name">
          {props.data}
          <div className="close-icon" />
        </div>
      </div>
      <MonacoDiffEditor
        height="400px"
        width="100%"
        original={JSON.stringify(findSchema(refList, props.leftSpec?.doc) || undefined, null, '\t')}
        value={JSON.stringify(findSchema(refList, props.rightSpec?.doc) || undefined, null, '\t')}
        options={{
          minimap: {
            enabled: false,
          },
          overviewRulerLanes: 0,
        }}
        editorDidMount={(diffEditor) => {
          if (!refs) return;
          monacoMount(diffEditor, refs, props.setActiveRefs, refName);
        }}
      />
    </div>
  );
}
