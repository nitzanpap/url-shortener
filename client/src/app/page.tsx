"use client"
import { UrlShortener } from "@/components/urlShortener/urlShortener"
import { generalStrings } from "@/constants/constants"
import { ShortUrlContext } from "@/hooks/useShortUrlContext"
import { Button } from "@mantine/core"
import { useState } from "react"
import { toast } from "react-toastify"
import styles from "./page.module.scss"

export default function Home() {
  const [shortUrl, setShortUrl] = useState<string>("")
  return (
    <section className={styles.pageContainer}>
      <header className={styles.header}></header>
      <main className={styles.main}>
        <section className={styles.titleContainer}>
          <h1 className={styles.title}>{generalStrings.title}</h1>
        </section>
        <section className={styles.contentContainer}>
          <ShortUrlContext.Provider value={{ shortUrl, setShortUrl }}>
            <UrlShortener />
            {shortUrl && (
              <>
                <section className={styles.shortUrlContainer}>
                  <a href={shortUrl} className={styles.shortUrl}>
                    {shortUrl}
                  </a>
                </section>
                <Button
                  className={styles.copyButton}
                  onClick={() => {
                    navigator.clipboard.writeText(shortUrl)
                    toast.success("Copied to clipboard")
                  }}
                >
                  Copy to clipboard
                </Button>
              </>
            )}
          </ShortUrlContext.Provider>
        </section>
      </main>
      <footer className={styles.footer}></footer>
    </section>
  )
}
