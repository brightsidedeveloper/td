import { Enemy } from './Enemy'
import { Tower } from './Tower'

export type Cell = {
  key: string
  x: number
  y: number
  isPath: boolean
  enemies: Enemy[]
  tower: Tower
}
