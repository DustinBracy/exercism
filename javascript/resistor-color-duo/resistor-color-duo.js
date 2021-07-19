import { colorCode } from "../resistor-color/resistor-color";

export const decodedValue = ([band1, band2]) => {
  return colorCode(band1) * 10 + colorCode(band2)
};

