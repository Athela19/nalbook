const baseURL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export type ApiError = { code: string; message: string };

export async function apiGet<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${baseURL}${path}`, {
    ...init,
    headers: {
      "Content-Type": "application/json",
      ...(init?.headers ?? {}),
    },
    cache: "no-store",
  });
  if (!res.ok) {
    const body = (await res.json().catch(() => null)) as ApiError | null;
    throw new Error(body?.message ?? `Permintaan gagal (${res.status})`);
  }
  return (await res.json()) as T;
}

export async function getBackendHealth(): Promise<{ status: string }> {
  return apiGet<{ status: string }>("/api/v1/health");
}
