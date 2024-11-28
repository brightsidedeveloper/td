import { useEffect, useRef } from 'react'
import { GameState, useGameStore } from './useGameStore'
import request from '../api/request'

export default function useGameWebSocket() {
  const { setAll } = useGameStore()

  const intervalRef = useRef<ReturnType<typeof setInterval>>()

  useEffect(() => {
    clearInterval(intervalRef.current)

    const ws = new WebSocket(import.meta.env.VITE_WEBSOCKET_URL)
    console.log(import.meta.env.VITE_WEBSOCKET_URL)
    ws.onopen = (e) => console.log('Opening web socket.', e)

    ws.onmessage = (e) => {
      const gameState = JSON.parse(e.data) as GameState
      console.log(gameState)

      setAll(gameState)
    }

    ws.onclose = () => {
      console.log('Closing web socket. Falling back to pulling')
      intervalRef.current = setInterval(() => request('/api/game/state', { method: 'GET' }).then(setAll), 300)
    }

    return () => {
      ws.close()
    }
  }, [])
}
