import { useNavigate } from "react-router";
import { Button } from "./Button";
import { useCosto } from "../App";

const formatearCosto = (costo) => {
  if (costo === 0) return "1 330 000 a 1 588 000";
  if (costo === 1) return "1 588 000 a 3 376 000";
  if (costo === 2) return "3 376 000 a 5 064 000";
  if (costo === 3) return "5 064 000 a 6 752 000";
};

export const Result = () => {
  const navigator = useNavigate();
  const { costo, setCosto } = useCosto();

  const handleHome = () => {
    setCosto("");
    navigator("/");
  };

  const handleCalculation = () => {
    setCosto("");
    navigator("/calculation");
  };

  return (
    <div className="result">
      <h1>Coste de Proyectos de Agua y Saneamiento en el Perú</h1>

      <div className="result__container">
        <h2>El costo aproximado del proyecto es</h2>
        <h1>
          {costo ? (
            <>
              <span>PEN</span> {formatearCosto(parseInt(costo))}
            </>
          ) : (
            <span>calculando...</span>
          )}
        </h1>

        <div className="result__container-btns">
          <Button btnText="Inicio" variant="outline" onClick={handleHome} />
          <Button btnText="Calcular" onClick={handleCalculation} />
        </div>
      </div>
    </div>
  );
};
