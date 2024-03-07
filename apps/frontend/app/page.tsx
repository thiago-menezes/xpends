"use client";
import { CreateSpend, SpendList } from "@xpends/features";
import { Header } from "@xpends/components";

export default function Home() {
  return (
    <main>
      <Header />
      <hr />
      <CreateSpend />
      <hr />
      <SpendList />
    </main>
  );
}
