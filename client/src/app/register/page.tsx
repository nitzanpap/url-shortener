'use client'

import { RegisterForm } from '@/components/RegisterForm'
import styles from '../login/page.module.scss'

export default function RegisterPage() {
  return (
    <main className={styles.main}>
      <div className={styles.container}>
        <h1>Register</h1>
        <RegisterForm />
      </div>
    </main>
  )
}