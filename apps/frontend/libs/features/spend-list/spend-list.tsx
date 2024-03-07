import { useGetSpendList } from "./hooks/useGetSpendList";

export const SpendList = () => {
  const { data: spends } = useGetSpendList();

  return (
    <div>
      {spends?.data.map((spend) => (
        <div key={spend.id}>
          <h2>{spend.name}</h2>
          <h3>{spend.category}</h3>
          <p>{spend.dueDate}</p>
          <p>{spend.price}</p>
          <p>{spend.done}</p>
        </div>
      ))}
    </div>
  );
};
