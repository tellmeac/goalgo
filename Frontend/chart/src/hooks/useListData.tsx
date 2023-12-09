import { useEffect, useState } from "react";
import {useGetListDataQuery} from '../query/useGetListDataQuery'
import { DataModel } from "../types/DataModel";
export const useListData = (): DataModel[] => {
  const result = useGetListDataQuery();
  const [data,setData] = useState<DataModel[]>([])
  useEffect(()=>{
    if (result.isFetched && !!result.data?.stamps.length) {
      setData(result.data.stamps)
    }
  },[result.isFetched])

  return data;
}