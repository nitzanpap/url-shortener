// Purpose: Next.js configuration file.
import path from "node:path"

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  compiler: {
    removeConsole: process.env.NODE_ENV !== "development",
  },
  sassOptions: {
    includePaths: [path.join(new URL(".", import.meta.url).pathname, "styles")],
  },
}

export default nextConfig
