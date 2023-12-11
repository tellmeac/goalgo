import React, { useEffect } from "react";
import ChartComponent from "react-apexcharts";
import { useListData } from "../hooks/useListData";
import { useNewData } from "../hooks/useNewData";

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


  const closePriseData = data.slice(startinterval, endinterval).map(ds =>
  ({
    x: getDate(ds.x),
    y: ds.y.close,
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

  const needPointData = data.slice(startinterval, endinterval).map(ds => ds.needPoint)

  const newData = useNewData();

  useEffect(() => {
    if (newData.length > 0) {
      candleStickData.push({
        x: getDate(newData[0].x),
        y: [newData[0].y.open, newData[0].y.high, newData[0].y.low, newData[0].y.close],
      });

      closePriseData.push({
        x: getDate(newData[0].x),
        y: newData[0].y.close,
      })

      topLineData.push({
        x: getDate(newData[0].x),
        y: newData[0].topLine,
      })

      downLineData.push({
        x: getDate(newData[0].x),
        y: newData[0].downLine,
      })

      blueLineData.push({
        x: getDate(newData[0].x),
        y: newData[0].blueLine,
      })

      needPointData.push(newData[0].needPoint)

      xaxisCategories.push( new Intl.DateTimeFormat('eu-RU',
      {
        timeZone: 'Europe/Moscow',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit'
      })
      .format(newData[0].x * 1000))
    }
  }, [newData])

  const pointsData: ApexDiscretePoint[] = [];
  needPointData.forEach(ds => {
    if (ds === true) {
      pointsData.push({
        seriesIndex: 4,
        dataPointIndex: needPointData.indexOf(ds),
        fillColor: '#10af10',
        strokeColor: '#10af10',
        size: 10,
      })
    }
    if (ds === false) {
      pointsData.push({
        seriesIndex: 4,
        dataPointIndex: needPointData.indexOf(ds),
        fillColor: '#FF0000',
        strokeColor: '#FF0000',
        size: 10,
      })
    }
  }
  );

  var options = {
    xaxis: {
      overwriteCategories: xaxisCategories,
      tickAmount: 30,
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
    colors: ['#000000', '#FF0000', '#00FF00', '#0000FF', '#00000000', '#000000'],
    markers: {
      discrete: pointsData
    }
  };

  const series = [
    {
      name: "Candle stick",
      type: "candlestick",
      data: candleStickData,
      zIndex: 11,
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
    },
    {
      name: "Reccomendation",
      type: "line",
      data: closePriseData,
      zIndex: 110,
    },
    {
      name: "Close",
      type: "line",
      data: closePriseData,
      zIndex: 110,
    }
  ]

  return <ChartComponent options={options} series={series} />
}
