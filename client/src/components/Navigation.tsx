import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
import styles from "./Navigation.module.scss";

export function Navigation() {
  const { isAuthenticated, logout } = useAuth();

  return (
    <nav className={styles.nav}>
      {/* ... existing navigation items ... */}

      {isAuthenticated ? (
        <button onClick={logout} className={styles.authButton}>
          Logout
        </button>
      ) : (
        <Link href="/login" className={styles.authButton}>
          Login
        </Link>
      )}
    </nav>
  );
}
