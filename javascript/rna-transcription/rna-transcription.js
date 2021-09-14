const DNA_RNA_MAP = {
  G: "C",
  C: "G",
  T: "A",
  A: "U",
};

export const toRna = (dna) => {
  return dna.replace(/[GCTA]/g, (strand) => DNA_RNA_MAP[strand]);
};
