import { createContext, useContext, useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import { About } from "./components/About";
import { Calculation } from "./components/Calculation";
import { Header } from "./components/Header";
import { Home } from "./components/Home";
import { MLP } from "./components/MLP";
import { Result } from "./components/Result";

const CostoContext = createContext();

export const useCosto = () => useContext(CostoContext);

function App() {
  const [costo, setCosto] = useState("");

  return (
    <div className="App">
      <CostoContext.Provider value={{ costo, setCosto }}>
        <Router>
          <Header />
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/calculation" element={<Calculation />} />
            <Route path="/result" element={<Result />} />
            <Route path="/about" element={<About />} />
            <Route path="/mlp" element={<MLP />} />
          </Routes>
        </Router>
      </CostoContext.Provider>
    </div>
  );
}

export default App;
