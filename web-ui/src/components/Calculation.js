import { useState } from "react";
import { useNavigate } from "react-router";
import { Button } from "./Button";
import { useCosto } from "../App";

export const Calculation = () => {
  const [nBeneficiarios, setNBeneficiarios] = useState("");
  const [nPuestos, setNPuestos] = useState("");
  const [region, setRegion] = useState("");
  const navigator = useNavigate();

  const { setCosto } = useCosto();

  const handleClick = (e) => {
    e.preventDefault();
    (async () => {
      const res = await fetch("http://localhost:9000/leer_proyecto", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ nBeneficiarios, nPuestos, region }),
      });
      const data = await res.json();
      setCosto(data?.costo);
    })();
    navigator("/result");
  };

  return (
    <div className="calculation">
      <h1>Coste de Proyectos de Agua y Saneamiento en el Perú</h1>

      <div className="calculation__container">
        <form>
          <div className="calculation__form-div">
            <label htmlFor="nBeneficiarios">Nº de beneficiarios:</label>
            <input
              id="nBeneficiarios"
              type="text"
              value={nBeneficiarios}
              onChange={(e) => setNBeneficiarios(e.target.value)}
            />
          </div>
          <div className="calculation__form-div">
            <label htmlFor="nPuestos">
              Nº de puestos de trabajo generados:
            </label>
            <input
              id="nPuestos"
              type="text"
              value={nPuestos}
              onChange={(e) => setNPuestos(e.target.value)}
            />
          </div>
          <div className="calculation__form-div">
            <label htmlFor="region">Region:</label>
            <input
              id="region"
              type="text"
              value={region}
              onChange={(e) => setRegion(e.target.value)}
            />
          </div>
          <Button btnText="Calcular" onClick={handleClick} />
        </form>
      </div>
    </div>
  );
};
