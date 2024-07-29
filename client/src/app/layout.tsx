import { configurations } from "@/configs/config"
import { generalStrings } from "@/constants/constants"
import type { Metadata } from "next"
import { Inter } from "next/font/google"
// All packages except `@mantine/hooks` require styles imports
import "@mantine/core/styles.css"
import "./global.scss"

import ToastProvider from "@/components/ToastProvider/ToastProvider"
import { ColorSchemeScript, createTheme, MantineProvider } from "@mantine/core"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: generalStrings.title,
  description: generalStrings.description,
  authors: [{ name: generalStrings.author.name, url: generalStrings.author.url }],
}

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
})

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  console.log(configurations.envVars.serverBaseUrl)

  return (
    <html lang="en">
      <head>
        <ColorSchemeScript />
      </head>
      <body className={inter.className}>
        <MantineProvider defaultColorScheme="dark" theme={theme}>
          <ToastProvider>{children}</ToastProvider>
        </MantineProvider>
      </body>
    </html>
  )
}
