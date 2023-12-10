import styled from 'styled-components';
import './App.css';
import { Chart } from './components/Chart';
import { Table } from './components/Table';

function App() {

  const Header = styled.div`
    font-size: 30px;
    color: white;
    font-weight: 500;
    background-color: #01c401;
    width: 100vw;
    height: 50px;
    font-family: Impact, Haettenschweiler, 'Arial Narrow Bold', sans-serif
  `

  const SubHeader = styled.div`
    font-size: 30px;
    color: white;
    font-weight: 500;
    background-color: #01c401;
    width: 100vw;
    height: 50px;
    font-family: Impact, Haettenschweiler, 'Arial Narrow Bold', sans-serif
  `

  const Span = styled.span`
    display: inline-flex;
    padding-top: 5px;
  `

  const ChartWrapper = styled.div`
    width: 75vw;
    display: inline-grid;
  `

  const Body = styled.div`
    display: inline-flex;
  `

  // TODO: добавить контекст для связи с таблицей. Настроить таблицу (сейчас она как заглушка).
  return (
    <div className="App">
      <Header>
        <Span>
          StraightCats
        </Span>
      </Header>
      <Body>
        <Table/>
        <ChartWrapper>
          <Chart />
        </ChartWrapper>
      </Body>
      <div>
        {'Telegram bot: '}
        <a href='https://t.me/demo_goalgo_publisher'>t.me/demo_goalgo_publisher</a>
      </div>
    </div>
  );
}

export default App;
