import axios, { AxiosResponse } from 'axios';
import { DataModel } from '../types/DataModel';

export const newData = (): Promise<AxiosResponse<DataModel[]>> => {
  const dateNow = Date.now();
  return axios.get(`http://localhost:8080/api/updates?from=${dateNow.toString()}`);
}