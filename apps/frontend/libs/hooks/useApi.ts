import axios from "axios";

export const useApi = () => {
  const getApi = () => {
    const baseURL = process.env["NEXT_PUBLIC_API"];
    const api = axios.create({
      baseURL,
    });

    api.interceptors.request.use(async (config) => {
      return config;
    });

    return api;
  };

  const api = getApi();

  return { api };
};
