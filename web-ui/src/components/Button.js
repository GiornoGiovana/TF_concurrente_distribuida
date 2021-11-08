import { useState } from "react";

export const Button = ({ btnText, variant = "normal", ...props }) => {
  const [className, _] = useState(variant);

  return (
    <button className={className} {...props}>
      {btnText}
    </button>
  );
};
