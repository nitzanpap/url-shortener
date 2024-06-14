"use client" // This is temporary until I extract the client code into a separate components

import { useEffect, useState } from "react"
import styles from "./page.module.scss"
import { generalStrings } from "@/constants/constants"
import { generateShortUrl, isServerAvailable } from "../api/serverApi"
import { isValidUrl } from "@/utils/utils"

export default function Home() {
  const [url, setUrl] = useState<string>("")

  const handleUrlInputChanged = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUrl(e.target.value)
  }

  const handleGenerateButtonClicked = async () => {
    // check if the URL is valid

    if (!isValidUrl(url)) {
      console.log("Invalid URL")
      return
    }

    // send the URL to the server
    const generatedUrl = await generateShortUrl(url)
    if (generatedUrl) {
      console.log("Generated URL:", generatedUrl)
    } else {
      console.log("Failed to generate short URL")
    }
  }

  useEffect(() => {
    new Promise<void>(async (resolve, reject) => {
      console.log("Is server available?", await isServerAvailable())
    })
  }, [])

  return (
    <section className={styles.pageContainer}>
      <header className={styles.header}></header>
      <main className={styles.main}>
        <section className={styles.titleContainer}>
          <h1 className={styles.title}>{generalStrings.title}</h1>
        </section>
        <section className={styles.contentContainer}>
          <input
            type="url"
            name="url"
            id="url"
            className={styles.urlInput}
            placeholder="Enter URL"
            value={url}
            onChange={handleUrlInputChanged}
          />
          <button className={styles.urlButton} onClick={handleGenerateButtonClicked}>
            Generate
          </button>
        </section>
      </main>
      <footer className={styles.footer}></footer>
    </section>
  )
}
