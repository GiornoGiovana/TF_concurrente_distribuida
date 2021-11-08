import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { Calculation } from "./components/Calculation";
import { Header } from "./components/Header";
import { Home } from "./components/Home";
import { Result } from "./components/Result";

function App() {
  return (
    <div className="App">
      <Router>
        <Header />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/calculation" element={<Calculation />} />
          <Route path="/result" element={<Result />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
