export const steps = (n) => {  
  if (n<1) throw Error("Only positive numbers are allowed")
  let counter = 0
  while (n > 1) {
    n = n % 2 === 0 ? n / 2 : n * 3 + 1
    counter++
  }
  return counter
};


