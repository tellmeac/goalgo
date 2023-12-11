import { AxiosResponse } from "axios";

export type ApiResponse<T> = AxiosResponse<{
  data: T;
  errorMessage: string;
  isSuccessful: boolean;
}>;