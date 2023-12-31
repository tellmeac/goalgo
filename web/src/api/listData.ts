import axios, { AxiosResponse } from 'axios';
import { DataModel } from '../types/DataModel';

export const listData = (): Promise<AxiosResponse<DataModel[]>> => {
//   const startTime = Date.now() - 2 * 24 * 60 * 60 * 1000 - demoDelay;
//   return axios.get(`http://localhost:8080/api/updates?from=${parseInt((startTime/1000).toString())}`);
  return axios.get(`http://localhost:8080/api/updates?from=${1701826800}`);
}