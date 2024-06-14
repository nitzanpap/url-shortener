import { UrlShortener } from "@/components/urlShortener/urlShortener"
import styles from "./page.module.scss"
import { generalStrings } from "@/constants/constants"

export default function Home() {
  return (
    <section className={styles.pageContainer}>
      <header className={styles.header}></header>
      <main className={styles.main}>
        <section className={styles.titleContainer}>
          <h1 className={styles.title}>{generalStrings.title}</h1>
        </section>
        <section className={styles.contentContainer}>
          <UrlShortener />
        </section>
      </main>
      <footer className={styles.footer}></footer>
    </section>
  )
}
