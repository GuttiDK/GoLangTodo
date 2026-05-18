import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Go Todo',
  description: 'Full-stack todo app with Go backend',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
