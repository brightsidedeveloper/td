import { create } from 'zustand'
import { Endpoints } from '../api/request'

export const useGameStore = create<GameStore>((set) => {
  return {
    grid: null,
    playerHealth: null,
    isRunning: null,
    setAll: (state) => set(state),
    clearAll: () => set({ grid: null, playerHealth: null, isRunning: null }),
    updateGrid: (grid) => set(() => ({ grid })),
  }
})

type GameStore = (GameState | BlankGameState) & {
  setAll: (state: GameState) => void
  clearAll: () => void
  updateGrid: (grid: GameState['grid']) => void
}

export type GameState = Endpoints['/api/game/state']

type BlankGameState = { grid: null; playerHealth: null; isRunning: null }
