import {Table as TableComponent} from 'antd'
import styled from 'styled-components';

export const Table = () => {
  const columns = [
    {
      title: 'Issuer of the shares',
      dataIndex: 'name',
      key: 'name',
    }
  ];

  const dataSource = [
    {
      key: '1',
      name: 'Sber',
    },
    {
      key: '2',
      name: 'RDSN',
    },
    {
      key: '2',
      name: 'LKOH',
    },
    {
      key: '2',
      name: 'NYTK',
    },
    {
      key: '2',
      name: 'SIBN',
    },
    {
      key: '2',
      name: 'GAZP',
    },
  ];

  const TableWrapper = styled(TableComponent)`
    width: 200px;
    border: 2px solid black;
    border-radius: 4px;
    height: 400px;
    position: relative;
    left: 0;

    thead > tr > th{
        background-color: #4fd43e !important;
        font-size: 16px;
      }

      tbody > tr:first-child > td {
        height: 52px;
        background-color: #4fd43e52 !important;
        border: 2px solid #4fd43e;
        border-radius: 4px ;
        font-size: 16px;
      }
  `;

  return <TableWrapper
            columns={columns}
            dataSource={dataSource}
            pagination={false}
        />
}