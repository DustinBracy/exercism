import { colorCode } from "../resistor-color/resistor-color";

export const decodedValue = (colors) => {
  let colorDigits = colors.slice(0,2).map((color) => colorCode(color))
  return (parseInt(colorDigits.join('')))
};
