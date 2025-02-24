import * as React from 'react';
import { Dialog, Message } from '@b-design/ui';
import type { ApplicationCompareResponse } from '../../interface/application';
import './index.less';
import { DiffEditor } from '../DiffEditor';
import { If } from 'tsx-control-statements/components';
import Translation from '../Translation';

type ApplicationDiffProps = {
  baseName: string;
  targetName: string;
  compare: ApplicationCompareResponse;
  onClose: () => void;
  id?: string;
};

export const ApplicationDiff = (props: ApplicationDiffProps) => {
  const { baseName, targetName, compare } = props;
  const container = 'revision' + baseName + targetName;
  return (
    <Dialog
      className={'commonDialog application-diff'}
      isFullScreen={true}
      footer={<div />}
      id={props.id}
      visible={true}
      onClose={props.onClose}
      title={
        <div style={{ color: '#fff' }}>
          {' Differences between '}
          <span className="name">{props.baseName}</span>
          {' and '}
          <span className="name">{props.targetName}</span>
        </div>
      }
    >
      <If condition={!compare.isDiff}>
        <Message type="success" style={{ marginBottom: '8px' }}>
          <Translation>There is no change</Translation>
        </Message>
      </If>
      <div id={container} className="diff-box">
        <DiffEditor
          language={'yaml'}
          id={container}
          target={compare.targetAppYAML}
          base={compare.baseAppYAML}
        />
      </div>
    </Dialog>
  );
};
