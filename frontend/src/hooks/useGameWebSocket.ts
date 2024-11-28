import { useEffect } from 'react'
import { GameState, useGameStore } from './useGameStore'

export default function useGameWebSocket() {
  const { setAll } = useGameStore()

  useEffect(() => {
    const ws = new WebSocket(import.meta.env.VITE_WEBSOCKET_URL)
    console.log(import.meta.env.VITE_WEBSOCKET_URL)
    ws.onopen = (e) => console.log('Opening web socket.', e)

    ws.onmessage = (e) => {
      const gameState = JSON.parse(e.data) as GameState
      console.log(gameState)

      setAll(gameState)
    }

    ws.onclose = () => console.log('Closing web socket.')

    return () => {
      ws.close()
    }
  }, [])
}
