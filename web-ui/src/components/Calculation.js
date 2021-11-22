import { useState } from "react";
import { useNavigate } from "react-router";
import { Button } from "./Button";
import { useCosto } from "../App";

const DEPARTAMENTOS = [
  "Amazonas",
  "Áncash",
  "Apurímac",
  "Arequipa",
  "Ayacucho",
  "Cajamarca",
  "Callao",
  "Cusco",
  "Huancavelica",
  "Huánuco",
  "Ica",
  "Junín",
  "La Libertad",
  "Lambayeque",
  "Lima",
  "Loreto",
  "Moquegua",
  "Pasco",
  "Piura",
  "Puno",
  "San Martín",
  "Tacna",
  "Tumbes",
  "Ucayali",
];

export const Calculation = () => {
  const [BENEFICIARIOS, setNBeneficiarios] = useState("");
  const [PUESTOS, setNPuestos] = useState("");
  const [REGION, setRegion] = useState(0);
  const [EPOCAS, setEpocas] = useState(20);
  const navigator = useNavigate();

  const { setCosto } = useCosto();

  const handleClick = (e) => {
    const COSTO = "5";
    e.preventDefault();
    (async () => {
      const res = await fetch("http://localhost:9000/epocas", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ EPOCAS: parseInt(EPOCAS) }),
      });
      const data = await res.json();
      if (data) {
        console.log(`Epocas enviadas: ${EPOCAS}`);
      }
    })();
    (async () => {
      const res = await fetch("http://localhost:9000/leer_proyecto", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          BENEFICIARIOS,
          PUESTOS,
          REGION: REGION.toString(),
          COSTO,
        }),
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
        <div className="calculation__form-div">
          <label htmlFor="nBeneficiarios">Nº de beneficiarios:</label>
          <input
            id="nBeneficiarios"
            type="text"
            value={BENEFICIARIOS}
            onChange={(e) => setNBeneficiarios(e.target.value)}
          />
        </div>
        <div className="calculation__form-div">
          <label htmlFor="nPuestos">Nº de puestos de trabajo generados:</label>
          <input
            id="nPuestos"
            type="text"
            value={PUESTOS}
            onChange={(e) => setNPuestos(e.target.value)}
          />
        </div>
        <div className="calculation__form-div">
          <label htmlFor="region">Region:</label>
          <select
            id="region"
            value={REGION}
            onChange={(e) => setRegion(e.target.value)}
          >
            {DEPARTAMENTOS.map((depa, idx) => (
              <option key={idx} value={idx}>
                {depa}
              </option>
            ))}
          </select>
        </div>
      </div>
      <div className="calculation__container">
        <div className="calculation__form-epocas">
          <label htmlFor="epocas">Nº Epocas</label>
          <p style={{ color: "red", margin: "5px 0px" }}>Campo opcional</p>
          <input
            id="epocas"
            type="text"
            value={EPOCAS}
            onChange={(e) => setEpocas(e.target.value)}
          />
        </div>
      </div>

      <Button btnText="Calcular" onClick={handleClick} />
    </div>
  );
};
