"use client";

import { useEffect, useState } from "react";
import { authClient } from "@/lib/auth-client";

export default function Home() {
  const [result, setResult] = useState<any>(null);

  useEffect(() => {
    authClient.getSession().then(setResult).catch((e) => setResult({ error: String(e) }));
  }, []);

  return <pre>{JSON.stringify(result, null, 2)}</pre>;
}
