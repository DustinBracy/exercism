import { colorCode } from "../resistor-color/resistor-color";

export const decodedValue = (colors) => {
  const [band1, band2] = colors
  return colorCode(band1) * 10 + colorCode(band2)
};
