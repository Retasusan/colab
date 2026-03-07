"use client";

import { FormEvent, useEffect, useState } from "react";
import { authClient } from "@/lib/auth-client";

type Organization = {
  id: number;
  name: string;
  slug: string;
};

export default function Home() {
  const [session, setSession] = useState<any>(null);
  const [orgs, setOrgs] = useState<Organization[]>([]);
  const [name, setName] = useState("");
  const [slug, setSlug] = useState("");
  const [error, setError] = useState("");

  async function getAccessToken() {
    const result = await authClient.token();
    console.log("token result", result);

    const token =
      result?.data?.token ??
      result?.data?.accessToken ??
      result?.token;

    if (!token) {
      throw new Error("failed to get auth token");
    }

    return token;
  }

  async function loadOrgs() {
    const token = await getAccessToken();

    const res = await fetch("http://localhost:8080/api/orgs", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error("failed to fetch organizations");
    }

    const data = await res.json();
    setOrgs(data);
  }

  useEffect(() => {
    authClient.getSession().then(async (s) => {
      setSession(s);
      if (s?.data?.session || s?.data?.user) {
        await loadOrgs();
      }
    }).catch(() => setSession(null));
  }, []);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setError("");

    try {
      const token = await getAccessToken();

      const res = await fetch("http://localhost:8080/api/orgs", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ name, slug }),
      });

      if (!res.ok) {
        throw new Error(await res.text());
      }

      setName("");
      setSlug("");
      await loadOrgs();
    } catch (e) {
      setError(e instanceof Error ? e.message : "failed");
    }
  }

  return (
    <main style={{ padding: 24 }}>
      <h1>colab</h1>

      <button onClick={() => authClient.signIn.social({ provider: "google" })}>
        Sign in with Google
      </button>

      <pre style={{ marginTop: 16 }}>{JSON.stringify(session, null, 2)}</pre>

      <hr style={{ margin: "24px 0" }} />

      <h2>Create organization</h2>
      <form onSubmit={onSubmit} style={{ display: "grid", gap: 8, maxWidth: 320 }}>
        <input
          placeholder="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          placeholder="slug"
          value={slug}
          onChange={(e) => setSlug(e.target.value)}
        />
        <button type="submit">Create</button>
      </form>

      {error && <p>{error}</p>}

      <hr style={{ margin: "24px 0" }} />

      <h2>Organizations</h2>
      <ul>
        {orgs?.map((org) => (
          <li key={org.id}>
            {org.name} ({org.slug})
          </li>
        ))}
      </ul>
    </main>
  );
}
