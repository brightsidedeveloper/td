import { createFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import request from '../api/request'
import { useGameStore } from '../hooks/useGameStore'

export const Route = createFileRoute('/')({
  async loader() {
    return request('/api/game/state', { method: 'GET' })
  },
  component: HomeComponent,
})

function HomeComponent() {
  const initGameState = Route.useLoaderData()
  const { grid, setAll } = useGameStore()
  // useGameWebSocket()
  useEffect(() => setAll(initGameState), [initGameState, setAll])
  if (!grid) return 'Loading...'

  console.log(grid)

  return (
    <div
      className="grid"
      style={{
        gridTemplateColumns: `repeat(${grid.length}, 1fr)`,
      }}
    >
      {grid.map((row) => {
        return row.map((cell) => {
          return (
            <div>
              <span className="text-white">{cell.x}</span>
              <span className="text-white">{cell.y}</span>
            </div>
          )
        })
      })}
    </div>
  )
}
