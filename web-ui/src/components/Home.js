import { useNavigate } from "react-router";
import { Button } from "./Button";

export const Home = () => {
  const navigator = useNavigate();

  const handleClick = () => {
    navigator("/calculation");
  };

  return (
    <div className="home">
      <div className="home__hero">
        <h1>Coste de Proyectos de Agua y Saneamiento en el Perú</h1>
        <p>
          Esta aplicación ha sido diseñada para todos los peruanos y su derecho
          de maternerse informado.
        </p>
        <Button btnText="Comenzar" onClick={handleClick} />
      </div>

      <div className="home__imgs">
        <img src="/images/image_1.png" alt="image_1" />
        <img src="/images/image_2.png" alt="image_2" />
      </div>
    </div>
  );
};
