"use client";

import { LoginForm } from "@/components/LoginForm";
import styles from "./page.module.scss";

export default function LoginPage() {
  return (
    <main className={styles.main}>
      <div className={styles.container}>
        <h1>Login</h1>
        <LoginForm />
      </div>
    </main>
  );
}
