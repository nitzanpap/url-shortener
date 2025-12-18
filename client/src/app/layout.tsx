import { generalStrings } from "@/constants/constants";
import type { Metadata, Viewport } from "next";
import { Inter } from "next/font/google";
// All packages except `@mantine/hooks` require styles imports
import "@mantine/core/styles.css";
import "./global.scss";

import ToastProvider from "@/components/ToastProvider/ToastProvider";
import { ColorSchemeScript, createTheme, MantineProvider } from "@mantine/core";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: generalStrings.title,
  description: generalStrings.description,
  generator: "Next.js",
  manifest: "/manifest.json",
  keywords: ["nextjs", "next14", "pwa", "next-pwa"],
  authors: [
    { name: generalStrings.author.name, url: generalStrings.author.url },
  ],
  icons: [
    { rel: "apple-touch-icon", url: "icons/icon-128x128.png" },
    { rel: "icon", url: "icons/icon-128x128.png" },
  ],
};

export const viewport: Viewport = {
  minimumScale: 1,
  initialScale: 1,
  width: "device-width",
  themeColor: [{ media: "(prefers-color-scheme: dark)", color: "#fff" }], // Moved here
  viewportFit: "cover",
};

const theme = createTheme({
  primaryColor: "purple",
  colors: {
    purple: [
      "#f3edff",
      "#e0d7fa",
      "#beabf0",
      "#9a7ce6",
      "#7c56de",
      "#683dd9",
      "#5f2fd8",
      "#4f23c0",
      "#451eac",
      "#3a1899",
    ],
  },
});

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  // Setting the default color scheme to dark
  const colorScheme = "dark";

  return (
    <html lang="en" data-mantine-color-scheme={colorScheme}>
      <head>
        {/* This will set the data-mantine-color-scheme attribute before React hydrates */}
        <ColorSchemeScript defaultColorScheme={colorScheme} />
      </head>
      <body className={inter.className}>
        <MantineProvider defaultColorScheme={colorScheme} theme={theme}>
          <ToastProvider>{children}</ToastProvider>
        </MantineProvider>
      </body>
    </html>
  );
}
