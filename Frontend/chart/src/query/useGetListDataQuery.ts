import { UseQueryResult, useQuery } from "react-query";
import { listData } from "../api/listData";
import { DataModel } from "../types/DataModel";
import { AxiosResponse } from "axios";
import { InputDataType } from "../types/InputDataType";

const fetch = async () => {
  const dataPromise: AxiosResponse<DataModel[]> = await listData();

  if (dataPromise.request.status !== 200) {
    throw new Error(dataPromise.statusText);
  }

  return dataPromise.data;
}

export const useGetListDataQuery = (): UseQueryResult<InputDataType> => {
  return useQuery({
    queryKey: ['getListData'],
    queryFn: fetch,
  })
}
