import React, { useEffect, useState } from "react";
import ChartComponent from "react-apexcharts";
import axios from 'axios';
import { useListData } from "./hooks/useListData";
import { format } from "date-fns";

export const Chart = () => {
  const data = useListData();

const startinterval = data.length - 10000;
const endinterval = data.length;

const xaxisCategories = data.slice(startinterval, endinterval).map(ds =>
  new Intl.DateTimeFormat('eu-RU',
  {
    timeZone: 'Europe/Moscow',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
  .format(ds.x * 1000)
)

  const getDate = (date: number) =>
   new Intl.DateTimeFormat('eu-RU',
  {
    timeZone: 'Europe/Moscow',
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  })
  .format(date * 1000)

  const candleStickData = data.slice(startinterval, endinterval).map(ds =>
  ({
    x: getDate(ds.x),
    y: [ds.y.open, ds.y.high, ds.y.low, ds.y.close],
  })
  )

  const topLineData = data.slice(startinterval, endinterval).map(ds =>
  ({
    x: getDate(ds.x),
    y: ds.topLine
  })
  )

  const downLineData = data.slice(startinterval, endinterval).map(ds =>
  ({
    x: getDate(ds.x),
    y: ds.downLine
  })
  )

  const blueLineData = data.slice(startinterval, endinterval).map(ds =>
  ({
    x: getDate(ds.x),
    y: ds.blueLine
  })
  )
  const pointsData: ApexDiscretePoint[] = [];
  data.slice(startinterval, endinterval).forEach(ds => {
    if (ds.needPoint === true) {
      pointsData.push({
        seriesIndex: 3,
        dataPointIndex: data.indexOf(ds),
        fillColor: '#00FF00',
        strokeColor: '#00ff00',
        size: 5,
      })
    }
    if (ds.needPoint === false) {
      pointsData.push({
        seriesIndex: 3,
        dataPointIndex: data.indexOf(ds),
        fillColor: '#FF0000',
        strokeColor: '#FF0000',
        size: 5,
      })
    }
  }
  );

  var options = {
    xaxis: {
      overwriteCategories: xaxisCategories,
      tickAmount: 20,
    },
    yaxis: {
      tooltip: {
        enabled: true
      },
      decimalsInFloat: 2,
    },
    stroke: {
      dashArray: [0, 8, 8]
    },
    colors: ['#000000', '#FF0000', '#00FF00', '#0000FF', '#FF0000', '#0000FF'],
    markers: {
      discrete: pointsData
    }
  };

  const series = [
    {
      name: "Candle stick",
      type: "candlestick",
      data: candleStickData,
      zIndex: 0,
      fill: {
        pattern: {
          style: 'width: 20px'
        }
      },
    }
    ,
    {
      name: "Upper band",
      type: "line",
      data: topLineData,
      zIndex: 1
    }
    ,
    {
      name: "Lower band",
      type: "line",
      data: downLineData,
      zIndex: 1
    }
    ,
    {
      name: "Kernel smooth",
      type: "line",
      data: blueLineData,
      zIndex: 1
    }
  ]

  return <ChartComponent options={options} series={series} />
}