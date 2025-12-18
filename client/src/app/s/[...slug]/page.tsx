import { getOriginalUrl } from "@/api/serverShortUrlApi";
import { redirect } from "next/navigation";

export default async function ShortUrlPage({
  params,
}: {
  params: { slug: string[] };
}) {
  const shortUrlHash = params.slug[0];
  const originalUrlResData = await getOriginalUrl(shortUrlHash);
  if (!originalUrlResData) {
    console.log("Failed to get original URL");
    return;
  }
  const { originalUrl } = originalUrlResData;
  redirect(originalUrl);
}
