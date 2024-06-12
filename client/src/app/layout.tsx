import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./global.scss"
import { generalStrings } from "@/constants/constants"
import { configurations } from "@/configs/config"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: generalStrings.title,
  description: generalStrings.description,
  authors: [{ name: generalStrings.author.name, url: generalStrings.author.url }],
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  console.log(configurations.envVars.serverBaseUrl);
  
  return (
    <html lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
