import { Link } from "react-router-dom";

export const Header = () => {
  return (
    <div className="header">
      <Link to="/">
        <h1>Despierta Perú</h1>
      </Link>

      <div className="header-btns">
        <Link to="/about">Nosotros</Link>
        <Link to="/mlp">Modelo MLP</Link>
      </div>
    </div>
  );
};
