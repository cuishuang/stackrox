import React from 'react';
import { generatePath, Link } from 'react-router-dom';
import {
    Divider,
    Pagination,
    Progress,
    ProgressMeasureLocation,
    Toolbar,
    ToolbarContent,
    ToolbarItem,
    Tooltip,
} from '@patternfly/react-core';
import { Table, Tbody, Td, Th, Thead, Tr } from '@patternfly/react-table';

import TbodyUnified from 'Components/TableStateTemplates/TbodyUnified';
import { UseURLPaginationResult } from 'hooks/useURLPagination';
import { UseURLSortResult } from 'hooks/useURLSort';
import { ComplianceClusterOverallStats } from 'services/ComplianceCommon';
import { getDistanceStrictAsPhrase } from 'utils/dateUtils';
import { TableUIState } from 'utils/getTableUIState';

import { coverageClusterDetailsPath } from './compliance.coverage.routes';
import {
    calculateCompliancePercentage,
    getCompliancePfClassName,
    getStatusCounts,
} from './compliance.coverage.utils';
import ProfilesTableToggleGroup from './components/ProfilesTableToggleGroup';
import StatusCountIcon from './components/StatusCountIcon';

export type ProfileClustersTableProps = {
    currentDatetime: Date;
    pagination: UseURLPaginationResult;
    profileClustersResultsCount: number;
    profileName: string;
    tableState: TableUIState<ComplianceClusterOverallStats>;
    getSortParams: UseURLSortResult['getSortParams'];
};

function ProfileClustersTable({
    currentDatetime,
    pagination,
    profileClustersResultsCount,
    profileName,
    tableState,
    getSortParams,
}: ProfileClustersTableProps) {
    const { page, perPage, setPage, setPerPage } = pagination;

    return (
        <>
            <Toolbar>
                <ToolbarContent>
                    <ToolbarItem>
                        <ProfilesTableToggleGroup activeToggle="clusters" />
                    </ToolbarItem>
                    <ToolbarItem variant="pagination" align={{ default: 'alignRight' }}>
                        <Pagination
                            itemCount={profileClustersResultsCount}
                            page={page}
                            perPage={perPage}
                            onSetPage={(_, newPage) => setPage(newPage)}
                            onPerPageSelect={(_, newPerPage) => setPerPage(newPerPage)}
                        />
                    </ToolbarItem>
                </ToolbarContent>
            </Toolbar>
            <Divider />
            <Table>
                <Thead>
                    <Tr>
                        <Th sort={getSortParams('Cluster ID')}>Cluster</Th>
                        <Th>Last scanned</Th>
                        <Th>Fail status</Th>
                        <Th>Pass status</Th>
                        <Th>Other status</Th>
                        <Th width={30}>Compliance</Th>
                    </Tr>
                </Thead>
                <TbodyUnified
                    tableState={tableState}
                    colSpan={6}
                    errorProps={{
                        title: 'There was an error loading profile clusters',
                    }}
                    emptyProps={{
                        message: 'No results found',
                    }}
                    filteredEmptyProps={{
                        title: 'No clusters found',
                        message: 'Clear all filters and try again',
                    }}
                    renderer={({ data }) => (
                        <Tbody>
                            {data.map((clusterInfo) => {
                                const {
                                    cluster: { clusterId, clusterName },
                                    lastScanTime,
                                    checkStats,
                                } = clusterInfo;
                                const { passCount, failCount, otherCount, totalCount } =
                                    getStatusCounts(checkStats);
                                const passPercentage = calculateCompliancePercentage(
                                    passCount,
                                    totalCount
                                );
                                const progressBarId = `progress-bar-${clusterId}`;
                                const firstDiscoveredAsPhrase = getDistanceStrictAsPhrase(
                                    lastScanTime,
                                    currentDatetime
                                );

                                return (
                                    <Tr key={clusterId}>
                                        <Td dataLabel="Cluster">
                                            <Link
                                                to={generatePath(coverageClusterDetailsPath, {
                                                    clusterId,
                                                    profileName,
                                                })}
                                            >
                                                {clusterName}
                                            </Link>
                                        </Td>
                                        <Td dataLabel="Last scanned">{firstDiscoveredAsPhrase}</Td>
                                        <Td dataLabel="Fail status">
                                            <StatusCountIcon
                                                text="check"
                                                status="fail"
                                                count={failCount}
                                            />
                                        </Td>
                                        <Td dataLabel="Pass status">
                                            <StatusCountIcon
                                                text="check"
                                                status="pass"
                                                count={passCount}
                                            />
                                        </Td>
                                        <Td dataLabel="Other status">
                                            <StatusCountIcon
                                                text="check"
                                                status="other"
                                                count={otherCount}
                                            />
                                        </Td>
                                        <Td dataLabel="Compliance">
                                            <Progress
                                                id={progressBarId}
                                                value={passPercentage}
                                                measureLocation={ProgressMeasureLocation.outside}
                                                className={getCompliancePfClassName(passPercentage)}
                                                aria-label={`${clusterName} compliance percentage`}
                                            />
                                            <Tooltip
                                                content={
                                                    <div>
                                                        {`${passCount} / ${totalCount} checks are passing for this cluster`}
                                                    </div>
                                                }
                                                triggerRef={() =>
                                                    document.getElementById(
                                                        progressBarId
                                                    ) as HTMLButtonElement
                                                }
                                            />
                                        </Td>
                                    </Tr>
                                );
                            })}
                        </Tbody>
                    )}
                />
            </Table>
        </>
    );
}

export default ProfileClustersTable;
