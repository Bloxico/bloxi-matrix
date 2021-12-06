import React from 'react'
import { PageLayout, Text, Icon, Link } from '@filecoin/ui'
import { getButton } from '@/pages/tests'
import { SystemScore, TestStatus } from '@filecoin/types'
import styled from 'styled-components'

const TABS = ['overview', 'detailedView'] as const
type Tab = typeof TABS[number]

interface HeaderProps {
  activeTab: Tab
  score: SystemScore | TestStatus
  onTabChange: (Tab) => void
  pageName: string
}

export const SystemHeader: React.FC<HeaderProps> = ({
  activeTab,
  score,
  onTabChange,
  pageName,
}) => {
  return (
    <PageLayout.Header>
      <BreadCrumbs>
        <Link to="/">System</Link>
        <StyledText type={'text xl'} bold>
          / {pageName}
        </StyledText>
        {getButton(score)}
      </BreadCrumbs>
      <PageLayout.Tabs>
        <PageLayout.Tab
          onClick={() => {
            onTabChange('overview')
          }}
          active={activeTab === 'overview'}
        >
          <Icon name="book" size="small" />
          <Text>Overview</Text>
        </PageLayout.Tab>
        <PageLayout.Tab
          onClick={() => onTabChange('detailedView')}
          active={activeTab === 'detailedView'}
        >
          <Icon name="detailed_view" size="small" />
          <Text>Detailed view</Text>
        </PageLayout.Tab>
      </PageLayout.Tabs>
    </PageLayout.Header>
  )
}

const BreadCrumbs = styled.div``
const StyledText = styled(Text)`
  padding-right: 15px;
  padding-left: 5px;
`
