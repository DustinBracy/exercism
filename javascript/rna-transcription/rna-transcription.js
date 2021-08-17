const DNA = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

export const toRna = (input) => {
  let output = [];
  input.split("").map((strand) => {
    output.push(DNA[strand]);
  });
  return output.join("");
};
