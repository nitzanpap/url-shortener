import Link from 'next/link'
import { useAuth } from '@/contexts/AuthContext'
import styles from './Navigation.module.scss'

export function Navigation() {
  const { user, signOut } = useAuth()

  const handleLogout = async () => {
    await signOut()
  }

  return (
    <nav className={styles.nav}>
      {/* ... existing navigation items ... */}
      
      {user ? (
        <div className={styles.userSection}>
          <span>Welcome, {user.email}</span>
          <button onClick={handleLogout} className={styles.authButton}>
            Logout
          </button>
        </div>
      ) : (
        <div className={styles.authLinks}>
          <Link href="/login" className={styles.authButton}>
            Login
          </Link>
          <Link href="/register" className={styles.authButton}>
            Register
          </Link>
        </div>
      )}
    </nav>
  )
} 
