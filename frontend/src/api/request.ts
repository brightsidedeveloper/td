import { GameGrid } from '../types/GameGrid'

export default async function request<T extends keyof Endpoints>(
  endpoint: T,
  { method = 'GET', body, params }: RequestOptions
): Promise<Endpoints[T]> {
  const query = params ? `?${new URLSearchParams(params).toString()}` : ''

  const response = await fetch(endpoint + query, {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
    body: method !== 'GET' ? JSON.stringify(body) : undefined,
  })

  if (!response.ok) {
    throw new Error(`Failed to fetch ${endpoint}`)
  }

  return response.json()
}

type RequestOptions = {
  method?: 'POST' | 'GET' | 'PUT' | 'DELETE'
} & (
  | {
      body: Record<string, unknown>
      params?: never
    }
  | {
      body?: never
      params?: Record<string, string>
    }
)

export type Endpoints = {
  '/api/game/state': {
    grid: GameGrid
    playerHealth: number
    isRunning: boolean
  }
}
