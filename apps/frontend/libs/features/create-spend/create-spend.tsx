import { useCreateSpend } from "./hooks/useCreateSpend";

export const CreateSpend = () => {
  const { mutate } = useCreateSpend();

  return (
    <div>
      <button onClick={() => mutate()}>Criar novo gasto</button>
    </div>
  );
};
