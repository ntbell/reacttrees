import './App.css';
import logo from './logo.svg';
import { TestJSXComponent } from './TestJSXComponent';
import { TestTSXComponent } from './TestTSXComponent';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
      <body>
        <TestTSXComponent />
        <TestJSXComponent />
        <p>Some text</p>
      </body>
    </div>
  );
}

export default App;
