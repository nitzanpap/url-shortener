import styles from "./page.module.scss"
import { generalStrings } from "@/constants/constants"

export default function Home() {
  return (
    <main className={styles.main}>
      <header className={styles.header}>
        <h1 className={styles.title}>{generalStrings.title}</h1>
      </header>
      <section className={styles.section}>
        
      </section>
    </main>
  )
}
