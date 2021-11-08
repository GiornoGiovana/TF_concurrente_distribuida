import { useNavigate } from "react-router";
import { Button } from "./Button";

export const Calculation = () => {
  const navigator = useNavigate();

  const handleClick = (e) => {
    e.preventDefault();
    navigator("/result");
  };

  return (
    <div className="calculation">
      <h1>Coste de Proyectos de Agua y Saneamiento en el Perú</h1>

      <div className="calculation__container">
        <form>
          <div className="calculation__form-div">
            <label htmlFor="nBeneficiarios">Nº de beneficiarios:</label>
            <input id="nBeneficiarios" type="text" />
          </div>
          <div className="calculation__form-div">
            <label htmlFor="nPuestos">
              Nº de puestos de trabajo generados:
            </label>
            <input id="nPuestos" type="text" />
          </div>
          <div className="calculation__form-div">
            <label htmlFor="region">Region:</label>
            <input id="region" type="text" />
          </div>
          <div className="calculation__form-div">
            <label htmlFor="uEjecutora">Unidad ejecutora:</label>
            <input id="uEjecutora" type="text" />
          </div>

          <Button btnText="Calcular" onClick={handleClick} />
        </form>
      </div>
    </div>
  );
};
