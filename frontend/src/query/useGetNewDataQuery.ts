//import React, { useState } from 'react';
import { UseQueryResult, useQuery } from "react-query";
import { DataModel } from "../types/DataModel";
import { AxiosResponse } from "axios";
import { InputDataType } from "../types/InputDataType";
import { newData } from "../api/newData";

const fetch = async () => {
  const dataPromise: AxiosResponse<DataModel[]> = await newData();

  if (dataPromise.request.status !== 200) {
    throw new Error(dataPromise.statusText);
  }

  return dataPromise.data;
}

export const useGetNewDataQuery = (): UseQueryResult<InputDataType> => {
  //const [refetchInterval, setRefetchInterval] = useState(5000)
  return useQuery({
    queryKey: ['getListData'],
    queryFn: fetch,
    refetchInterval: 1,
  })
}
