import { useMutation } from "@tanstack/react-query";
import { useApi } from "@xpends/hooks";

export const useCreateSpend = () => {
  const { api } = useApi();
  const payload = {
    name: "Cartão de crédito",
    category: "banco",
    price: 4847,
    done: false,
    duedate: "07/03/2024",
  };

  const body = payload;

  return useMutation({
    mutationKey: ["create-spend"],
    mutationFn: () => api.post("/", body),
  });
};
