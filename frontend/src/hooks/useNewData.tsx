import { useEffect, useState } from "react";
import { useGetNewDataQuery } from '../query/useGetNewDataQuery'
import { DataModel } from "../types/DataModel";

export const useNewData = (): DataModel[] => {
  const result = useGetNewDataQuery();
  const [data,setData] = useState<DataModel[]>([])
  useEffect(()=>{
    if (result.isFetched && !!result.data?.stamps.length) {
      setData(result.data.stamps)
    }
  },[result.isFetched])

  return data;
}