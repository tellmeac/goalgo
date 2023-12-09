import styled from 'styled-components';
import './App.css';
import { Chart } from './Chart';

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

  const Span = styled.span`
    display: inline-flex;
    padding-top: 5px;
  `

  return (
    <div className="App">
      <Header>
        <Span>
          StraightCats
        </Span>
      </Header>
      <Chart/>
      <div>
      {'Telegram bot: '}
        <a href='t.me/demo_goalgo_publisher'>t.me/demo_goalgo_publisher</a>
      </div>
    </div>
  );
}

export default App;
