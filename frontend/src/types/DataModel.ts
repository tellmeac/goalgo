export type DataModel = 
  {
    x: number,
    y: {
      open: number,
      high: number,
      low: number,
      close: number,
    },
    topLine: number,
    downLine: number,
    blueLine: number,
    needPoint: boolean | undefined
  }