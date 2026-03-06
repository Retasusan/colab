"use client";

import { useEffect, useState } from "react";
import { authClient } from "@/lib/auth-client";

export default function Home() {
  const [session, setSession] = useState<any>(null);

  useEffect(() => {
    authClient.getSession().then(setSession).catch(() => setSession(null));
  }, []);

  return (
    <main style={{ padding: 24 }}>
      <h1>colab</h1>

      <button
        onClick={() => authClient.signIn.social({ provider: "google" })}
        style={{ padding: 12, border: "1px solid #ccc" }}
      >
        Sign in with Google
      </button>

      <pre style={{ marginTop: 16 }}>{JSON.stringify(session, null, 2)}</pre>
    </main>
  );
}