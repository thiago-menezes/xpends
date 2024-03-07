import { useQuery } from "@tanstack/react-query";
import { useApi } from "@xpends/hooks";
import { AxiosError } from "axios";

export const useGetSpendList = () => {
  const { api } = useApi();
  return useQuery<{ data: TSpendList }, AxiosError>({
    queryKey: ["spends"],
    queryFn: () => api.get("/"),
    retry: false,
  });
};
