import { createFileRoute } from '@tanstack/react-router'
import { useEffect } from 'react'
import request from '../api/request'
import { useGameStore } from '../hooks/useGameStore'
import clsx from 'clsx'
import useGameWebSocket from '../hooks/useGameWebSocket'

export const Route = createFileRoute('/')({
  async loader() {
    return request('/api/game/state', { method: 'GET' })
  },
  component: HomeComponent,
})

function HomeComponent() {
  const initGameState = Route.useLoaderData()
  const { grid, playerHealth, isRunning, setAll } = useGameStore()
  useGameWebSocket()
  useEffect(() => setAll(initGameState), [initGameState, setAll])
  if (!grid) return 'Loading...'

  console.log(playerHealth ?? 0)

  return (
    <div className="w-fit">
      <div className="relative h-4 bg-gray-400">
        <div className="size-full bg-green-500" style={{ width: `${playerHealth}%` }} />
      </div>
      <div className="flex justify-between">
        <button
          className="p-4 bg-emerald-500"
          onClick={() => {
            request('/api/game/start', { method: 'POST' }).catch((e) => console.log('Error starting game.', e.message))
          }}
        >
          Start
        </button>
        {isRunning && (
          <button
            onClick={() => {
              request('/api/game/reset', { method: 'POST' }).catch((e) => console.log('Error resetting game.', e.message))
            }}
          >
            Reset
          </button>
        )}
      </div>
      <div
        className="grid w-fit"
        style={{
          gridTemplateColumns: `repeat(${grid.length}, 1fr)`,
        }}
      >
        {grid.map((row) => {
          return row.map(({ key, x, y, isPath, tower, enemies }) => {
            return (
              <button
                key={key}
                onClick={() => {
                  console.log('Adding tower at: ', x, y)
                  request('/api/game/addTower', {
                    method: 'POST',
                    body: {
                      x,
                      y,
                    },
                  })
                    .then(() => console.log('Successfully added tower at: ', x, y))
                    .catch((e) => console.log('Error adding tower at: ', x, y, e.message))
                }}
                className={clsx(
                  'size-10',
                  enemies?.length ? 'bg-red-500' : isPath ? 'bg-green-500' : tower.isActive ? 'bg-blue-400' : 'bg-slate-400'
                )}
              />
            )
          })
        })}
      </div>
    </div>
  )
}
