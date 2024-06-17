export default function ShortUrlPageLayout({
  children,
  params,
}: {
  children: React.ReactNode
  params: { slug: string[] }
}) {
  return <>{children}</>
}
