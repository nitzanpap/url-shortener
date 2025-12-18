"use client"
import { UrlShortener } from "@/components/urlShortener/urlShortener"
import { generalStrings } from "@/constants/constants"
import { ShortUrlContext } from "@/hooks/useShortUrlContext"
import { ActionIcon } from "@mantine/core"
import { useEffect, useState } from "react"
import styles from "./page.module.scss"
import { ClipboardIcon } from "@/components/icons/ClipboardIcon"
import { Check } from "@/components/icons/Check"

export default function Home() {
  const [shortUrl, setShortUrl] = useState<string>("")
  const [copied, setCopied] = useState(false)

  useEffect(() => {
    if (copied) {
      navigator.clipboard.writeText(shortUrl)
    }
    const timeout = setTimeout(() => {
      if (copied) {
        setCopied(false)
      }
    }, 1000)

    return () => clearTimeout(timeout)
  }, [copied, shortUrl])
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
                  <ActionIcon
                    className={styles.copyButton}
                    data-copied={copied}
                    onClick={() => {
                      setCopied(true)
                    }}
                  >
                    <ClipboardIcon className={styles.clipboardIcon} />
                    <Check className={styles.checkIcon} />
                  </ActionIcon>
                </section>
              </>
            )}
          </ShortUrlContext.Provider>
        </section>
      </main>
      <footer className={styles.footer}></footer>
    </section>
  )
}
