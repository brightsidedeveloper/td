import { createFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import request from '../api/request'
import { useGameStore } from '../hooks/useGameStore'

export const Route = createFileRoute('/')({
  component: HomeComponent,
})

function HomeComponent() {
  const { grid, setAll } = useGameStore()

  useEffect(() => {
    request('/api/game/state', { method: 'GET' }).then((data) => {
      setAll(data)
    })
  }, [])

  if (!grid) return 'Loading...'

  return (
    <>
      {grid.map((row, i) => {
        row.map((cell, j) => {
          return
        })
      })}
    </>
  )
}
