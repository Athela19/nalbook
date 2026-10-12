import { getBackendHealth } from "../lib/api-client";

export const dynamic = "force-dynamic";

export default async function Home() {
  let backend = "tidak terhubung";
  try {
    const h = await getBackendHealth();
    backend = h.status;
  } catch {
    backend = "tidak terhubung";
  }
  return (
    <main style={{ padding: "var(--space-8)", maxWidth: 640 }}>
      <h1>Nalbook</h1>
      <p>Status backend: {backend}</p>
    </main>
  );
}
