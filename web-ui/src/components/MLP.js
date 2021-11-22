export const MLP = () => {
  return (
    <div
      style={{
        width: 750,
        marginTop: 16,
        margin: "auto",
        textAlign: "center",
      }}
    >
      <h1 style={{ color: "#545355" }}>Modelo de Machine Learning</h1>
      <h2 style={{ color: "#B5B4F9" }}>Multilayer Perceptron</h2>
      <p>
        Es una red neuronal de clase feedforward, la cual es mejor usada para
        casos cuando los datos de entrada son alta dimensión y discretos, y la
        salida es un valor real.
      </p>

      <img src="/images/mlp.png" alt="png" />
    </div>
  );
};
