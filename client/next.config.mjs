// Purpose: Next.js configuration file.
import path from "path"

// @ts-check

/** @type {import('next').NextConfig} */
const nextConfig = {
  sassOptions: {
    includePaths: [path.join(new URL(".", import.meta.url).pathname, "styles")],
  },
}

export default nextConfig
