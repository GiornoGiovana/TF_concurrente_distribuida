import { useNavigate } from "react-router";
import { Button } from "./Button";

export const Result = () => {
  const navigator = useNavigate();

  const handleHome = () => {
    navigator("/");
  };

  const handleCalculation = () => {
    navigator("/calculation");
  };

  return (
    <div className="result">
      <h1>Coste de Proyectos de Agua y Saneamiento en el Perú</h1>

      <div className="result__container">
        <h2>El costo aproximado del proyecto es</h2>
        <h1>
          <span>PEN</span> 7,486,458.15
        </h1>

        <div className="result__container-btns">
          <Button btnText="Inicio" variant="outline" onClick={handleHome} />
          <Button btnText="Calcular" onClick={handleCalculation} />
        </div>
      </div>
    </div>
  );
};
