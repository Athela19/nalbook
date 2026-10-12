import type { Metadata } from "next";
import "../styles/tokens.css";

export const metadata: Metadata = {
  title: "Nalbook",
  description: "Software akuntansi web untuk UMKM dan kantor akuntan Indonesia.",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="id">
      <body>{children}</body>
    </html>
  );
}
